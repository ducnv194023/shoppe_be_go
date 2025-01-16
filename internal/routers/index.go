package routers

import (
	user "github.com/ducnv194023/shoppe_be_go/internal/routers/user"
	manager "github.com/ducnv194023/shoppe_be_go/internal/routers/manager"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)