package admin

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"net/http"
	"strconv"
)

type PhotoController struct {
	BaseController
}

func (pc *PhotoController) PostList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	rows, _ := strconv.Atoi(ctx.PostForm("rows"))
	limit_start := (page - 1) * rows

	search := ctx.PostForm("search")
	search_str := ""
	if search != "" { //topjui prompt 有BUG
		search_str = "`title` like \"%" + search + "%\""
	}

	type photoM struct {
		model.Photo
		TTitle  string `json:"t_title"`
		Creater string `json:"creater"`
	}

	list := []photoM{}
	if search != "" {
		service.InitDB().Where(search_str).Offset(limit_start).Limit(rows).Find(&list)
	} else {
		service.InitDB().Offset(limit_start).Limit(rows).Find(&list)
	}
	var count int64
	if search != "" {
		service.InitDB().Model(model.Photo{}).Where(search_str).Count(&count)
	} else {
		service.InitDB().Model(model.Photo{}).Count(&count)
	}
	var pages int
	if len(list)%rows == 0 {
		pages = len(list) / rows
	} else {
		pages = len(list)/rows + 1
	}

	//var list2 []articleM
	for i, vo := range list {
		if vo.Tid != "" {
			categoryM := model.Category{}
			service.InitDB().Model(categoryM).Where(map[string]interface{}{"classid": vo.Tid}).Find(&categoryM)
			//vo.TTitle = categoryM.Classtitle
			//使用(&list[i])不用引入第三方变量
			(&list[i]).TTitle = categoryM.Classtitle
			userM := model.User{}
			service.InitDB().Model(userM).Where(map[string]interface{}{"id": vo.CreaterId}).Find(&userM)
			(&list[i]).Creater = userM.Nickname
		}
		//list2 = append(list2, vo)
	}
	//fmt.Printf("%+v", list2)
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": pages, "total": count, "rows": list})
}

func (pc *PhotoController) PostZree(ctx *gin.Context) {
	classpid, _ := strconv.Atoi(ctx.Query("classpid"))
	category_lst := new(service.Category).ZreeList("1", classpid, map[string]interface{}{"classmodule": "photo"})
	ctx.JSON(http.StatusOK, category_lst)
}

func (pc *PhotoController) PostAdd(ctx *gin.Context) {
	photoM := model.Photo{}
	photoM.Tid = ctx.PostForm("tid")
	photoM.Title = ctx.PostForm("title")
	photoM.Img = ctx.PostForm("img")
	photoM.CreaterId = uint(AuthId)
	photoM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	res := service.InitDB().Create(&photoM)
	if res.Error != nil {
		pc.TopjuiError(ctx, res.Error.Error())
	} else {
		pc.TopjuiSucess(ctx, "新增成功")
	}
}

func (pc *PhotoController) GetEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		pc.TopjuiError(ctx, "参数Id丢失")
	}
	photoM := model.Photo{}
	service.InitDB().Where(map[string]interface{}{"id": id}).Find(&photoM)
	ViewData["info"] = photoM
	pc.BaseController.GetEdit(ctx)
}

func (pc *PhotoController) PostEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		pc.TopjuiError(ctx, "参数Id丢失")
	}
	photoM := model.Photo{}
	photoM.Tid = ctx.PostForm("tid")
	photoM.Title = ctx.PostForm("title")
	photoM.Img = ctx.PostForm("img")
	photoM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	res := service.InitDB().Where(map[string]interface{}{"id": id}).Updates(&photoM)
	if res.Error != nil {
		pc.TopjuiError(ctx, res.Error.Error())
	} else {
		pc.TopjuiSucess(ctx, "保存成功")
	}
}
