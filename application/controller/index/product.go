package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	BaseController
}

func (pc *ProductController) GetIndex(ctx *gin.Context) {
	ViewData["nav"] = "product_index"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/index.html", ViewData)
}
