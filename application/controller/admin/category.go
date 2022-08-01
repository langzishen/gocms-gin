package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"net/http"
	"strconv"
	"strings"
)

type CategoryController struct {
	BaseController
}

func (cc *CategoryController) PostList(ctx *gin.Context) {
	classpid, _ := strconv.Atoi(ctx.Query("classpid"))
	category_lst := new(service.Category).ZreeList("*", classpid)
	ctx.JSON(http.StatusOK, category_lst)
}

func (cc *CategoryController) GetAdd(ctx *gin.Context) {
	model_M := new(service.Model).GetModelList()
	var model_list []map[string]interface{}
	for _, model_vo := range model_M {
		model_list = append(model_list, map[string]interface{}{"text": model_vo.Cname, "value": model_vo.Ename})
	}
	model_byte, _ := json.Marshal(model_list)
	ViewData["model_list"] = string(model_byte)
	cc.BaseController.GetAdd(ctx)
}

func (cc *CategoryController) PostZree(ctx *gin.Context) {
	classpid, _ := strconv.Atoi(ctx.Query("classpid"))
	category_lst_first := []map[string]interface{}{{"classid": 0, "classtitle": "顶级分类", "state": "open"}}
	category_lst := new(service.Category).ZreeList("*", classpid)
	if classpid == 0 {
		category_lst = append(category_lst_first, category_lst...)
	}
	ctx.JSON(http.StatusOK, category_lst)
}

func (cc *CategoryController) PostAdd(ctx *gin.Context) {
	categoryM := model.Category{}
	categoryM.Classpid, _ = strconv.Atoi(ctx.PostForm("classpid"))
	categoryM.Classchild = 0
	categoryM.Classtitle = ctx.PostForm("classtitle")
	categoryM.Classsort = 999
	categoryM.Classstatus, _ = strconv.Atoi(ctx.PostForm("classstatus"))
	categoryM.Classkeywords = ctx.PostForm("classkeywords")
	categoryM.Classdescription = ctx.PostForm("classdescription")
	categoryM.Classmodule = ctx.PostForm("classmodule")
	categoryM.Classimg = ctx.PostForm("classimg")
	categoryM.Classapv = 0
	categoryM.Classmenushow, _ = strconv.Atoi(ctx.PostForm("classmenushow"))
	service.InitDB().Create(&categoryM)
	/**
	if categoryM.Classpid != 0 {
		model.InitDB().Model(model.Category{}).Where(map[string]interface{}{"classid": categoryM.Classpid}).Update("classchild", 1)
	}**/
	cc.TopjuiSucess(ctx, "新增成功")
}

func (cc *CategoryController) GetEdit(ctx *gin.Context) {
	classid, _ := strconv.Atoi(ctx.Query("classid"))
	if classid == 0 {
		cc.TopjuiError(ctx, "参数Id丢失")
	}
	categoryM := model.Category{}
	service.InitDB().Where(map[string]interface{}{"classid": classid}).Find(&categoryM)
	model_M := new(service.Model).GetModelList()
	var model_list []map[string]interface{}
	for _, model_vo := range model_M {
		if model_vo.Ename == categoryM.Classmodule {
			model_list = append(model_list, map[string]interface{}{"text": model_vo.Cname, "value": model_vo.Ename, "selected": true})
		} else {
			model_list = append(model_list, map[string]interface{}{"text": model_vo.Cname, "value": model_vo.Ename})
		}
	}
	model_byte, _ := json.Marshal(model_list)
	ViewData["model_list"] = string(model_byte)
	ViewData["info"] = categoryM
	cc.BaseController.GetEdit(ctx)
}

func (cc *CategoryController) PostEdit(ctx *gin.Context) {
	classid, _ := strconv.Atoi(ctx.PostForm("classid"))
	if classid == 0 {
		cc.TopjuiError(ctx, "参数Id丢失")
	}
	categoryM := model.Category{}
	categoryM.Classpid, _ = strconv.Atoi(ctx.PostForm("classpid"))
	categoryM.Classtitle = ctx.PostForm("classtitle")
	categoryM.Classstatus, _ = strconv.Atoi(ctx.PostForm("classstatus"))
	categoryM.Classkeywords = ctx.PostForm("classkeywords")
	categoryM.Classdescription = ctx.PostForm("classdescription")
	categoryM.Classmodule = ctx.PostForm("classmodule")
	categoryM.Classimg = ctx.PostForm("classimg")
	categoryM.Classmenushow, _ = strconv.Atoi(ctx.PostForm("classmenushow"))
	res := service.InitDB().Where(map[string]interface{}{"classid": classid}).Updates(categoryM)
	if res.Error != nil {
		cc.TopjuiError(ctx, res.Error.Error())
	}
	/**if categoryM.Classpid != 0 {
		model.InitDB().Model(model.Category{}).Where(map[string]interface{}{"classid": categoryM.Classpid}).Update("classchild", 1)
	}**/
	cc.TopjuiSucess(ctx, "修改成功")
}

func (cc *CategoryController) PostDel(ctx *gin.Context) {
	cc.AjaxDone(ctx, "del")
}

func (cc *CategoryController) PostResume(ctx *gin.Context) {
	cc.AjaxDone(ctx, "resume")
}

func (cc *CategoryController) PostForbid(ctx *gin.Context) {
	cc.AjaxDone(ctx, "forbid")
}

func (cc *CategoryController) PostRecycle(ctx *gin.Context) {
	cc.AjaxDone(ctx, "recycle")
}

func (cc *CategoryController) PostForever_del(ctx *gin.Context) {
	id := ctx.PostForm("classid")
	if id == "" {
		cc.TopjuiError(ctx, "删除操作参数Id必须")
	}
	var id_arr []int
	for _, id_str := range strings.Split(id, ",") {
		id_i, _ := strconv.Atoi(id_str)
		id_arr = append(id_arr, id_i)
	}
	res := service.InitDB().Exec("DELETE from "+RequestController+" WHERE classid IN ?", id_arr)
	if res.Error == nil {
		cc.TopjuiSucess(ctx, "删除成功")
	} else {
		cc.TopjuiError(ctx, "删除失败")
	}
}

func (cc *CategoryController) AjaxDone(ctx *gin.Context, do string) {
	id := ctx.PostForm("classid")
	if id == "" {
		cc.TopjuiError(ctx, "删除操作参数Id必须")
	}
	var status int
	var say string
	switch do {
	case "del": //删除      1->-1
		status = -1
		say = "删除"
	case "resume": //审核   0->1
		status = 1
		say = "审核"
	case "forbid": //禁用   1->0
		status = 0
		say = "禁用"
	case "recycle": //恢复   -1->0
		status = 1
		say = "恢复"
	}
	var id_arr []int
	for _, id_str := range strings.Split(id, ",") {
		id_i, _ := strconv.Atoi(id_str)
		id_arr = append(id_arr, id_i)
	}

	res := service.InitDB().Exec("UPDATE "+RequestController+" SET classstatus="+strconv.Itoa(status)+" WHERE classid IN ?", id_arr)
	if res.Error == nil {
		cc.TopjuiSucess(ctx, say+"成功")
	} else {
		cc.TopjuiError(ctx, say+"失败")
	}
}
