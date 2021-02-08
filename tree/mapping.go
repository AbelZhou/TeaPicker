package tree

import (
	"TeaPicker/store"
	"crypto/md5"
	"github.com/bluele/gcache"
	"github.com/json-iterator/go"
	"time"
)

/*
1、实现多个树并存 			done
2、实现一个简单的缓存系统？
*/
//创建mapping cacheSize缓存文章数量512 cacheExpire缓存时间（秒)3600
func CreateMapping(cacheSize int, cacheExpire int) *Mapping {
	m := Mapping{}
	m.mapping = make(map[string]*node)
	m.index_count = make(map[string]int)

	if cacheExpire == 0 {
		cacheExpire = 3600
	}
	if cacheSize == 0 {
		cacheSize = 256
	}

	m.readCache = gcache.New(cacheSize).
		Expiration(time.Second * time.Duration(cacheExpire)).
		ARC().
		Build()
	return &m
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//增加一个词
func (m *Mapping) AddWord(index string, key string, jsonData string, chinese string) bool {
	m.Lock()
	if m.mapping[index] == nil {
		m.mapping[index] = &node{}
	}

	node := m.mapping[index]
	res := node.AddWords(key, jsonData, chinese)
	if res {
		m.index_count[index] += 1
	}
	m.Unlock()
	return res
}

//删除一个词
func (m *Mapping) DelWord(index string, key string, chinese string) bool {
	m.Lock()
	if m.mapping[index] == nil {
		m.Unlock()
		return false
	}

	res := m.mapping[index].DelWords(key, chinese)
	if res {
		m.index_count[index] -= 1
	}
	m.Unlock()
	return res
}

//搜索
func (m *Mapping) Search(index string, str string) map[string]returnRes {
	//构造缓存key
	indexStringByte := []byte(index + str)
	cacheKey := md5.Sum(indexStringByte)
	//1000字(中文)以上加入缓存前置控制
	if len(str) > 3000 {
		cacheReturnRes, err := m.readCache.Get(cacheKey)
		if err == nil && cacheReturnRes != nil {
			return cacheReturnRes.(map[string]returnRes)
		}
	}

	//执行查询逻辑
	m.RLock()
	if m.mapping[index] == nil {
		m.RUnlock()
		return nil
	}
	res := m.mapping[index].Search(str)
	m.RUnlock()
	//加入缓存后置控制
	//if len(res) > 0 {
	if len(str) > 3000 {
		m.readCache.Set(cacheKey, res)
	}
	//}
	return res
}

func (m *Mapping) SearchLenovos(index string, str string) map[string]returnRes  {
	m.RLock()
	if m.mapping[index] == nil {
		m.RUnlock()
		return nil
	}
	res := m.mapping[index].SearchLenovos(str)
	m.RUnlock()
	return res

}

//获得索引内的词量
func (m *Mapping) getIndexCount(index string) int {
	m.RLock()
	if m.mapping[index] == nil {
		m.RUnlock()
		return 0
	}
	res := m.index_count[index]
	m.RUnlock()
	return res
}

//还原一条记录
func (m *Mapping) RecoveryLog(opera store.OperateType, commandString string) {
	switch opera {
	case store.OperateAdd:
		var addCommand AddWordCommand
		err := json.UnmarshalFromString(commandString, &addCommand)
		if err == nil {
			m.AddWord(addCommand.Index, addCommand.Key, addCommand.JsonData, addCommand.Chinese)
		}
		break
	case store.OperateDel:
		var delCommand DelWordCommand
		err := json.UnmarshalFromString(commandString, &delCommand)
		if err == nil {
			m.DelWord(delCommand.Index, delCommand.Key, delCommand.Chinese)
		}
		break
	}
}
//优化AOF日志
func (m *Mapping) Optimize(command store.Command, tuningMap store.TuningMap) string {
	var ret_key string
	if command.Operate == store.OperateDel {
		var delCommand DelWordCommand
		err := json.UnmarshalFromString(command.CommandContext, &delCommand)
		if err == nil {
			key := delCommand.Index+delCommand.Key
			if delCommand.Chinese != "" {
				key += delCommand.Chinese
			}
			if _, ok := tuningMap[key]; ok {
				delete(tuningMap, key)
			}
		}
		return ret_key
	}
	if command.Operate == store.OperateAdd {
		var addCommand AddWordCommand
		err := json.UnmarshalFromString(command.CommandContext, &addCommand)
		if err == nil {
			log, err := json.MarshalToString(store.Command{command.Time, command.Tag, command.Operate, command.CommandContext})
			if err != nil {
				return ret_key
			}
			key := addCommand.Index+addCommand.Key
			if addCommand.Chinese != "" {
				key += addCommand.Chinese
			}
			//直接覆盖 存在则保留后者
			tuningMap[key] = log
			return key
		}
	}
	return ret_key
}
