package content

import "github.com/gogf/gf/v2/frame/g"

type GetContentReq struct {
	g.Meta `path:"/get" tags:"content" method:"get" summary:"content"`
	Key    string `v:"required"`
}

type GetContentRes struct {
	Content string
}

type AddReq struct {
	g.Meta `path:"/add" tags:"content" method:"get" summary:"content"`
	Key    string `v:"required"`
	Type   string `v:"required"`
}

type AddRes struct {
}

type DeleteReq struct {
	g.Meta `path:"/delete" tags:"content" method:"get" summary:"content"`
	Key    string `v:"required"`
}

type DeleteRes struct {
}

type SaveContentReq struct {
	g.Meta  `path:"/save" tags:"content" method:"post" summary:"content"`
	Key     string `v:"required"`
	Content string
}

type SaveContentRes struct {
}

type DownloadImgReq struct {
	g.Meta `path:"/downloadImg" tags:"content" method:"post" summary:"content"`
	ImgUrl string `v:"required"`
}

type DownloadImgRes struct {
}
