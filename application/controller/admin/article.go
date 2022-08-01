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

type ArticleController struct {
	BaseController
}

func (ac *ArticleController) PostList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	rows, _ := strconv.Atoi(ctx.PostForm("rows"))
	limit_start := (page - 1) * rows

	search := ctx.PostForm("search")
	search_str := ""
	if search != "" { //topjui prompt 有BUG
		search_str = "`title` like \"%" + search + "%\""
	}

	type articleM struct {
		model.Article
		TTitle  string `json:"t_title"`
		Creater string `json:"creater"`
	}

	list := []articleM{}
	if search != "" {
		service.InitDB().Where(search_str).Offset(limit_start).Limit(rows).Find(&list)
	} else {

		service.InitDB().Offset(limit_start).Limit(rows).Find(&list)
	}
	var count int64
	if search != "" {
		service.InitDB().Model(model.Article{}).Where(search_str).Count(&count)
	} else {
		service.InitDB().Model(model.Article{}).Count(&count)
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

func (ac *ArticleController) GetAdd(ctx *gin.Context) {
	attrlist := new(service.Attribute).GetList(map[string]interface{}{"model_ename": "article"})
	ViewData["attrlist"] = attrlist
	ac.BaseController.GetAdd(ctx)
}

func (ac *ArticleController) PostZree(ctx *gin.Context) {
	classpid, _ := strconv.Atoi(ctx.Query("classpid"))
	category_lst := new(service.Category).ZreeList("1", classpid, map[string]interface{}{"classmodule": "article"})
	ctx.JSON(http.StatusOK, category_lst)
}

func (ac *ArticleController) PostAdd(ctx *gin.Context) {
	articleM := model.Article{}
	articleM.Tid = ctx.PostForm("tid")
	articleM.Title = ctx.PostForm("title")
	articleM.Keywords = ctx.PostForm("keywords")
	articleM.Description = ctx.PostForm("description")
	articleM.Img = ctx.PostForm("img")
	articleM.Content = ctx.PostForm("content")
	articleM.CreaterId = AuthId
	articleM.Sort, _ = strconv.Atoi(ctx.PostForm("sort"))
	articleM.Apv = 0
	articleM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	Attrtj_str, _ := ctx.GetPostFormArray("attrtj")
	articleM.Attrtj = strings.Join(Attrtj_str, ",")
	IsLoginShow, _ := strconv.Atoi(ctx.PostForm("is_login_show"))
	articleM.IsLoginShow = uint(IsLoginShow)
	res := service.InitDB().Create(&articleM)
	if res.Error != nil {
		ac.TopjuiError(ctx, res.Error.Error())
	} else {
		ac.TopjuiSucess(ctx, "新增成功")
	}
}

func (ac *ArticleController) GetEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		ac.TopjuiError(ctx, "参数Id丢失")
	}
	articleM := model.Article{}
	service.InitDB().Where(map[string]interface{}{"id": id}).Find(&articleM)

	attrtj := strings.Split(articleM.Attrtj, ",")
	attrlist := new(service.Attribute).GetList(map[string]interface{}{"model_ename": "article"})
	var attrlist2 []map[string]interface{}
	if len(attrlist) > 0 {
		for _, vo := range attrlist {
			map_byte, _ := json.Marshal(vo)
			map_vo := map[string]interface{}{}
			json.Unmarshal(map_byte, &map_vo)
			map_vo["is_check"] = IsInArr(attrtj, vo.Attrvalue)
			attrlist2 = append(attrlist2, map_vo)
		}
	}
	ViewData["attrlist"] = attrlist2
	ViewData["info"] = articleM
	ac.BaseController.GetEdit(ctx)
}

func (ac *ArticleController) PostEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		ac.TopjuiError(ctx, "参数Id丢失")
	}
	articleM := model.Article{}
	articleM.Tid = ctx.PostForm("tid")
	articleM.Title = ctx.PostForm("title")
	articleM.Keywords = ctx.PostForm("keywords")
	articleM.Description = ctx.PostForm("description")
	articleM.Img = ctx.PostForm("img")
	articleM.Content = ctx.PostForm("content")
	articleM.Sort, _ = strconv.Atoi(ctx.PostForm("sort"))
	articleM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	Attrtj_str, _ := ctx.GetPostFormArray("attrtj")
	articleM.Attrtj = strings.Join(Attrtj_str, ",")
	IsLoginShow, _ := strconv.Atoi(ctx.PostForm("is_login_show"))
	articleM.IsLoginShow = uint(IsLoginShow)
	res := service.InitDB().Where(map[string]interface{}{"id": id}).Updates(&articleM)
	if res.Error != nil {
		ac.TopjuiError(ctx, res.Error.Error())
	} else {
		ac.TopjuiSucess(ctx, "修改成功")
	}
}

func (ac *ArticleController) GetInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	title := ctx.Query("title")
	info := map[string]interface{}{"id": id, "title": title}
	ViewData["info"] = info
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/info.html", ViewData)
}
