package tree

import (
	"sync"
	"github.com/bluele/gcache"
)


//返回结果
type returnRes struct {
	Str  string
	Data string
}

//节点
type node struct {
	ch   string          //该节点字符
	data string          //Json挂载数据 自定义节点数据
	str  string          //最终字符串    需要提取的字符串 rune切片
	strList  map[string]string //拼音可能对应的多个词
	end  bool            //是否最终节点  default false
	list map[int32]*node //子节点       下级node rune表示后均为int32
}

//特征map
type Mapping struct {
	index_count map[string]int
	mapping     map[string]*node
	readCache   gcache.Cache
	sync.RWMutex
}

//store 生成command以及还原command枚举&结构
type AddWordCommand struct {
	Index    string
	Key      string
	JsonData string
	Chinese string
}

type DelWordCommand struct {
	Index string
	Key   string
	Chinese string
}
