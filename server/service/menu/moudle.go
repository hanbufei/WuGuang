package menu

import "fmt"

type Menu struct {
	Title  string `json:"title"`  //菜单标题，对应文件名或目录名
	Key    string `json:"key"`    //菜单值，对应路径
	IsLeaf bool   `json:"isLeaf"` //叶节点，对应是否为文件
}

func (m *Menu) toString() string {
	return fmt.Sprintf("{title:'%s',key:'%s',isLeaf:%t}", m.Title, m.Key, m.IsLeaf)
}
