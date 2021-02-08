package tree

import (
	"testing"
	"fmt"
)

func TestMapping_AddWord(t *testing.T) {
	m := CreateMapping()
	m.AddWord("link", "张三", "zhangsan")
	m.AddWord("link", "张三李", "zhangsanli")
	c := m.getIndexCount("link")
	//m.DelWord("link","张三")
	c = m.getIndexCount("link")
	fmt.Println(c)
	searchRes := m.Search("link", "张三李")
	fmt.Println(searchRes)
}
