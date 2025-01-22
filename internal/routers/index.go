package routers

import (
	auth "github.com/ducnv194023/shoppe_be_go/internal/routers/auth"
)

type RouterGroup struct {
	Auth auth.AuthRouterGroup
}

var RouterGroupApp = new(RouterGroup)