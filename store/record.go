// author:Abel
// email:abel.zhou@hotmail.com
// 2018/10/18

package store

import (
	"bufio"
	"fmt"
	"github.com/json-iterator/go"
	"io"
	"os"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Record struct {
	path      string "data.tmd"
	f         *os.File
	tags      map[TagType]IStore
	TuningMap TuningMap
	//buff chan Command
}

func CreateRecord(dataPath string) *Record {
	r := Record{}
	r.path = dataPath
	r.tags = make(map[TagType]IStore, 10)
	//r.buff = make(chan Command, 1000)
	openFile, err := os.OpenFile(r.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	r.f = openFile
	r.TuningMap = make(map[string]string)
	return &r

}

//接收写入请求
func (recd *Record) AddLog(tag TagType, operate OperateType, argsString string) {
	log, err := json.MarshalToString(Command{time.Now(), tag, operate, argsString})
	if err == nil {
		recd.f.WriteString(log + "\n")
	}
}

func (recd *Record) RegisterStoreStruct(tag TagType, t IStore) {
	recd.tags[tag] = t
}

func (recd *Record) Recovery() {
	var command Command
	readBuff := bufio.NewReader(recd.f)
	count := 0
	for {
		lineString, err := readBuff.ReadBytes('\n')
		if err == io.EOF {
			fmt.Printf("共执行[%d]条记录\n", count)
			return
		}
		err = json.Unmarshal(lineString, &command)
		if err != nil {
			continue
		}
		if _, ok := recd.tags[command.Tag]; ok {
			recd.tags[command.Tag].RecoveryLog(command.Operate, command.CommandContext)
			count++
		} else {
			//找不到已注册的对应结构
		}
	}
}
func (recd *Record) cleanup() []string {
	var command Command
	var keySlice []string
	count := 0
	readBuff := bufio.NewReader(recd.f)
	for {
		lineString, err := readBuff.ReadBytes('\n')
		if err == io.EOF {
			fmt.Printf("共清理%d条数\n", count - len(recd.TuningMap) )
			count = count - len(recd.TuningMap)
			return keySlice
		}
		err = json.Unmarshal(lineString, &command)
		if err != nil {
			continue
		}
		if _, ok := recd.tags[command.Tag]; ok {
			key := recd.tags[command.Tag].Optimize(command, recd.TuningMap)
			keySlice = append(keySlice, key)
			count++
		}
	}
}
func (recd *Record) Optimize() {
	fmt.Println("优化数据文件...")
	keySlice := recd.cleanup()
	var openFile *os.File
	openFile, err := os.OpenFile(recd.path+".tmp", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err == nil {
		for _, keys := range keySlice{
			if _, ok := recd.TuningMap[keys]; ok{
				openFile.WriteString(recd.TuningMap[keys] + "\n")
				delete(recd.TuningMap, keys)
			}
		}
		recd.f.Close()
		openFile.Close()
		os.Rename(recd.path+".tmp", recd.path)
	}
	fmt.Println("优化数据文件完毕")
}
