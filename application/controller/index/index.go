package index

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"net/http"
)

type IndexController struct {
	BaseController
}

func (ic *IndexController) GetIndex(ctx *gin.Context) {
	ViewData["nav"] = "index_index"

	categoryM := model.Category{}
	service.InitDB().Where(map[string]interface{}{"classmodule": "photo", "classtitle": "banner"}).Find(&categoryM)
	banner_list := []model.Photo{}
	service.InitDB().Where(map[string]interface{}{"tid": categoryM.Classid}).Find(&banner_list)
	ViewData["banner_list"] = banner_list

	article_list := []model.Article{}
	service.InitDB().Where("attrtj like '%1%'").Find(&article_list)
	ViewData["article_list"] = article_list

	ctx.HTML(http.StatusOK, "index/"+RequestController+"/index.html", ViewData)
}
