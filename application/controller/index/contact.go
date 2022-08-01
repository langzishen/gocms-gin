package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactController struct {
	BaseController
}

func (cc *ContactController) GetIndex(ctx *gin.Context) {
	ViewData["nav"] = "contact_index"
	ViewData["index"] = "hello word"
	ctx.HTML(http.StatusOK, "index/"+RequestController+"/index.html", ViewData)
}
