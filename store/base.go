// author:Abel
// email:abel.zhou@hotmail.com
// 2018/10/19

package store

import "time"

//结构标签
type TagType int8

const (
	TAG_Tree TagType = 0
)

//操作指令
type Command struct {
	Time           time.Time
	Tag            TagType
	Operate        OperateType
	CommandContext string
}

type TuningMap map[string]string

//操作类别
type OperateType int8

const (
	OperateAdd OperateType = 0
	OperateDel OperateType = 1
)

type IStore interface {
	RecoveryLog(opera OperateType, command string)
	Optimize(command Command, TuningMap TuningMap) string
}
