package version

import "github.com/gogf/gf/v2/frame/g"

type Req struct {
	g.Meta `path:"/" tags:"version" method:"get" summary:"version"`
}
type Res struct {
	g.Meta `mime:"text/html" example:"string"`
}
