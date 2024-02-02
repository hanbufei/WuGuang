package storage

import "github.com/gogf/gf/v2/frame/g"

type GetRootPathReq struct {
	g.Meta `path:"/getRootPath" tags:"storage" method:"get" summary:"storage"`
}
type GetRootPathRes struct {
	g.Meta   `mime:"application/json"`
	RootPath string
}

type SetRootPathReq struct {
	g.Meta   `path:"/setRootPath" tags:"storage" method:"get" summary:"storage"`
	RootPath string `v:"required"`
}

type SetRootPathRes struct {
	g.Meta `mime:"application/json"`
}
