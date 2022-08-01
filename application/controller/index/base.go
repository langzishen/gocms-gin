package index

import (
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	RequestApp        string //访问的当前app
	RequestController string //访问的当前控制器
	RequestAction     string //访问的当前操作
)
var ViewData = make(gin.H)

type BaseController struct {
}

/**
初始化
*/
func initialize(ctx *gin.Context) {
	ViewData["RequestApp"] = RequestApp
	ViewData["RequestController"] = RequestController
	ViewData["RequestAction"] = RequestAction

}

func IndexAuth(ctx *gin.Context) {
	full_path := ctx.FullPath()
	pathArr := strings.Split(full_path, "/")
	RequestApp = pathArr[1]        //app
	RequestController = pathArr[2] //controller
	RequestAction = pathArr[3]     //action
	initialize(ctx)
}
