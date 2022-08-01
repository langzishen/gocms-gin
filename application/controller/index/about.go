package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AboutController struct {
	BaseController
}

func (ac *AboutController) GetIndex(ctx *gin.Context) {
	ViewData["nav"] = "about_index"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/index.html", ViewData)
}

func (ac *AboutController) GetJoin(ctx *gin.Context) {
	ViewData["nav"] = "about_join"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/join.html", ViewData)
}

func (ac *AboutController) GetJoininfo(ctx *gin.Context) {
	ViewData["nav"] = "about_join"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/join_info.html", ViewData)
}
