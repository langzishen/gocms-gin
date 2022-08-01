package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
	BaseController
}

func (ac *ArticleController) GetIndex(ctx *gin.Context) {
	ViewData["nav"] = "article_index"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/index.html", ViewData)
}
