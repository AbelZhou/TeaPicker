package tree

import (
    "strings"
)

/**
实现tire-trees算法，具备增删检索功能 done
 */

//增加词
func (n *node) AddWords(str string, json_data string, chinese string) bool {
    //检查当前node list是否为未初始化的map
    if (n.list == nil) {
        n.list = make(map[int32]*node)
    }
    runeStr := []rune(str)
    //游标
    var curr_node *node
    curr_map := n.list
    for k, v := range runeStr {
        if curr_map[v] != nil {
            curr_node = curr_map[v]
        } else {
            //赋值
            curr_node = &node{
                ch:   string(v),
                list: make(map[int32]*node),
                strList:make(map[string]string),
            }

            curr_map[v] = curr_node
        }

        if (k == (len(runeStr) - 1)) {
            //最后一个字符 重新定义节点
            curr_node.end = true
            curr_node.data = json_data
            curr_node.str = str
            if chinese != "" {
                curr_node.strList[chinese] = json_data
            }
            //fmt.Println(curr_node.strList)
        } else {
            //非最后一个字符 重新定位map继续向下跑
            curr_map = curr_node.list
        }

    }
    return true
}

//删除词
func (n *node) DelWords(str string, chinese string) bool {
    if (n.list == nil) {
        return true;
    }
    runeStr := []rune(str)

    var pre_node *node
    var del_index int32
    curr_map := n.list

    //0字符
    if (len(runeStr) == 0) {
        return false
    }

    //只有一个字符直接清掉
    if (len(runeStr) == 1) {
        if (curr_map[runeStr[0]] != nil && len(curr_map[runeStr[0]].list) == 0) {
            delete(curr_map, runeStr[0])
            return true
        }
    }

    for k, v := range runeStr {
        if (curr_map[v] == nil) {
            //检查节点脱靶
            return false
        }

        //检查最终的叶节点
        if (k == (len(runeStr) - 1 )  ) {
            //要清理的字符节点不是叶
            if (curr_map[v].end == false) {
                //没找到字符串
                return false;
            }
            if chinese != "" {
        	    if _, ok := curr_map[v].strList[chinese]; ok {
        	        delete(curr_map[v].strList, chinese)
        	        if len(curr_map[v].strList) > 0 {
        	            return true;
                    }
                }
            }
            //是叶 检查子集个数
            if (len(curr_map[v].list) > 0) {
                //存在子集 直接自杀 不影响子集
                curr_map[v].end = false
                curr_map[v].data = ""
                curr_map[v].str = ""
                return true
            } else {
                //不存在子集 清理上游分叉节点
                if (pre_node != nil && del_index != 0) {
                    delete(pre_node.list, del_index)
                    return true
                } else {
                    //上游在根节点
                    delete(n.list, runeStr[0])
                    return true
                }
            }
        }
        //上游存在分叉
        if (curr_map[v].end == true || len(curr_map[v].list) > 1) {
            //记录需要删除的node以及需要删除的点
            pre_node = curr_map[v]
            del_index = runeStr[k+1]
        }
        curr_map = curr_map[v].list
    }
    return false
}

//搜索词
func (n *node) Search(context string) map[string]returnRes {
    context = strings.TrimSpace(context)
    //定义返回值
    res_map := make(map[string]returnRes)
    if len(context) == 0 {
        return res_map
    }

    context_rune := []rune(context)

    c := 0
    //协程在手
    for k := 0; k < len(context_rune); k++ {
        //回收结果
        res_channel := make(chan []returnRes)
        c++
        go n.scan(context_rune[k:], res_channel)
        result, ok := <-res_channel
        if (ok) {
            for _, res := range result {
                key := string(res.Str)
                res_map[key] = res
            }
        }
    }
    return res_map
}

//检索
func (n *node) scan(str []rune, res_channel chan []returnRes) {
    var returnList []returnRes;
    curr_node := n
    for _, index := range str {
        //没命中直接返回 阻塞关闭通道
        if curr_node.list[index] == nil {
            defer close(res_channel)
            break
        }
        //命中游标向下
        curr_node = curr_node.list[index]
        if curr_node.end {
            //增加返回结果
            returnList = append(returnList, returnRes{curr_node.str, curr_node.data})
        }
    }
    res_channel <- returnList
}

//搜索联想词
func (n *node) SearchLenovos(context string) map[string]returnRes  {
    context = strings.TrimSpace(context)
    res_map := make(map[string]returnRes)
    if len(context) == 0 {
        return res_map
    }
    context_rune := []rune(context)
    curr_node := n

    for k, index := range context_rune {
        if curr_node.list[index] == nil {
            return res_map
        }
        curr_node = curr_node.list[index]
        if k == len(context_rune) - 1 {
            temp_node := n.extract(curr_node, []returnRes{})
            for _, result := range temp_node {
                key := string(result.Str)
                res_map[key] = result
            }
        }
    }
    return res_map
}

//递归提取树中的所有结果
func (n *node) extract(curr_node *node, returnList []returnRes) []returnRes {
    if curr_node.end == true && len(curr_node.strList) > 0{
        for chinese, result := range curr_node.strList {
            returnList = append(returnList, returnRes{chinese, result})
        }
    }
    if len(curr_node.list) > 0 && curr_node.end == false {
        for _, temp_node := range curr_node.list {
            returnList = n.extract(temp_node, returnList)
        }
    }
    return returnList
}

