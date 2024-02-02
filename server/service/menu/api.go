package menu

import (
	"github.com/gogf/gf/v2/frame/g"
)

type InitMenuReq struct {
	g.Meta `path:"/init" tags:"menu" method:"get" summary:"menu"`
}

type InitMenuRes struct {
	g.Meta       `mime:"application/json"`
	MenuInitData []Menu
	Label        string
}

type ListMenuReq struct {
	g.Meta `path:"/list" tags:"menu" method:"get" summary:"menu"`
	Key    string
}

type ListMenuRes struct {
	g.Meta       `mime:"application/json"`
	MenuListData []Menu
}
