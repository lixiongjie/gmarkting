// file: controllers/user_controller.go

package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris/services"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SService
	Session *sessions.Session
}

func (c *IndexController) Get() mvc.Result {
	return nil
}
