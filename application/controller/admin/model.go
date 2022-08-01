package admin

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"net/http"
	"strconv"
)

type ModelController struct {
	BaseController
}

func (mc *ModelController) PostList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	rows, _ := strconv.Atoi(ctx.PostForm("rows"))
	limit_start := (page - 1) * rows

	ename := ctx.PostForm("ename")
	list := []model.Model{}
	var count int64
	if ename == "" {
		service.InitDB().Offset(limit_start).Limit(rows).Find(&list)
		service.InitDB().Model(model.Model{}).Count(&count)
	} else {
		service.InitDB().Where("ename like ?", "%"+ename+"%").Offset(limit_start).Limit(rows).Find(&list)
		service.InitDB().Model(model.Model{}).Where("ename like ?", "%"+ename+"%").Count(&count)
	}
	var pages int
	if len(list)%rows == 0 {
		pages = len(list) / rows
	} else {
		pages = len(list)/rows + 1
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": pages, "total": count, "rows": list})
}

func (mc *ModelController) PostAdd(ctx *gin.Context) {
	modelM := model.Model{}
	modelM.Ename = ctx.PostForm("ename")
	if modelM.Ename == "" {
		mc.TopjuiError(ctx, "模型名必填")
	} else {
		var count int64
		service.InitDB().Model(modelM).Where(map[string]interface{}{"ename": modelM.Ename}).Count(&count)
		if count > 0 {
			mc.TopjuiError(ctx, "模型名已经存在")
		}
	}
	modelM.Cname = ctx.PostForm("cname")
	modelM.Sort, _ = strconv.Atoi(ctx.PostForm("sort"))
	modelM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	res := service.InitDB().Model(modelM).Create(&modelM)
	if res.Error == nil {
		mc.TopjuiSucess(ctx, "新增成功")
	} else {
		mc.TopjuiError(ctx, "新增失败")
	}
}

func (mc *ModelController) GetEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		mc.TopjuiError(ctx, "Id参数丢失")
	}
	modelM := model.Model{}
	service.InitDB().Model(modelM).Where(map[string]interface{}{"id": id}).Find(&modelM)
	ViewData["info"] = modelM
	mc.BaseController.GetEdit(ctx)
}

func (mc *ModelController) PostEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		mc.TopjuiError(ctx, "Id参数丢失")
	}
	modelM := model.Model{}
	modelM.Ename = ctx.PostForm("ename")
	modelM.Cname = ctx.PostForm("cname")
	modelM.Sort, _ = strconv.Atoi(ctx.PostForm("sort"))
	modelM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	res := service.InitDB().Model(modelM).Where(map[string]interface{}{"id": id}).Updates(&modelM)
	if res.Error == nil {
		mc.TopjuiSucess(ctx, "编辑成功")
	} else {
		mc.TopjuiError(ctx, "编辑失败")
	}
}
