package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

/**
列表
*/
func (uc *UserController) PostList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	rows, _ := strconv.Atoi(ctx.PostForm("rows"))
	limit_start := (page - 1) * rows

	search := ctx.PostForm("search_value")
	search_str := "`id`>1" //排除admin超级管理员
	if search != "" {      //topjui prompt 有BUG
		search_str += " and `" + ctx.PostForm("search_key") + "` like \"%" + search + "%\""
	}

	type user struct {
		model.User
		LastLoginTime int    `json:"last_login_time"`
		LastLoginIp   string `json:"last_login_ip"`
		LoginCount    int64  `json:"login_count"`
	}

	list := []*user{}
	service.InitDB().Where(search_str).Offset(limit_start).Limit(rows).Find(&list)
	var count int64
	service.InitDB().Model(model.User{}).Where(search_str).Count(&count)
	var pages int
	if len(list)%rows == 0 {
		pages = len(list) / rows
	} else {
		pages = len(list)/rows + 1
	}

	for _, vo := range list {
		vo.LastLoginTime, vo.LastLoginIp, vo.LoginCount = new(service.LogLogin).LastLogLogin("boss", vo.Id)
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": pages, "total": count, "rows": list})
}

func (uc *UserController) PostAdd(ctx *gin.Context) {
	userM := model.User{}
	userM.ObjectId = RandString(8)
	userM.Account = ctx.PostForm("account")
	userM.Nickname = ctx.PostForm("nickname")
	userM.Password = fmt.Sprintf("%x", md5.Sum([]byte(ctx.PostForm("password"))))
	userM.Phone = ctx.PostForm("phone")
	userM.Email = ctx.PostForm("email")
	userM.CreateTime = uint(time.Now().Unix())
	userM.UpdateTime = uint(time.Now().Unix())
	userM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	userM.TypeId = 1
	res := service.InitDB().Create(&userM)
	if res.Error == nil {
		uc.TopjuiSucess(ctx, "新增成功")
	} else {
		uc.TopjuiError(ctx, "新增失败")
	}
}

func (uc *UserController) GetEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		uc.TopjuiError(ctx, "参数Id必须")
	}
	userM := model.User{}
	service.InitDB().Find(&userM, id)
	ViewData["info"] = userM
	uc.BaseController.GetEdit(ctx)
}

func (uc *UserController) PostEdit(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		uc.TopjuiError(ctx, "参数Id必须")
	}
	userM := model.User{}
	userM.Id = uint(id)
	userM.Account = ctx.PostForm("account")
	userM.Nickname = ctx.PostForm("nickname")
	userM.Phone = ctx.PostForm("phone")
	userM.Email = ctx.PostForm("email")
	userM.UpdateTime = uint(time.Now().Unix())
	userM.Status, _ = strconv.Atoi(ctx.PostForm("status"))
	res := service.InitDB().Save(&userM)
	if res.Error == nil {
		uc.TopjuiSucess(ctx, "保存成功")
	} else {
		uc.TopjuiError(ctx, "保存失败")
	}
}

func (uc *UserController) GetResetpassword(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		uc.TopjuiError(ctx, "参数Id必须")
	}
	ViewData["id"] = id
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/resetpassword.html", ViewData)
}

func (uc *UserController) PostResetpassword(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		uc.TopjuiError(ctx, "参数Id必须")
	}
	userM := model.User{}
	userM.Id = uint(id)
	userM.Password = fmt.Sprintf("%x", md5.Sum([]byte(ctx.PostForm("password"))))
	res := service.InitDB().Save(&userM)
	if res.Error == nil {
		uc.TopjuiSucess(ctx, "重置成功")
	} else {
		uc.TopjuiError(ctx, "重置失败")
	}
}

func (uc *UserController) GetPersonal(ctx *gin.Context) {
	userM := model.User{}
	service.InitDB().Where(map[string]interface{}{"id": AuthId}).Find(&userM)
	ViewData["info"] = userM
	role_user_list := []model.RoleUser{}
	service.InitDB().Where(map[string]interface{}{"id": AuthId}).Find(&role_user_list)
	role_name := []string{}
	if AuthTypeId == 9 {
		role_name = append(role_name, "超级管理员")
	}
	if len(role_user_list) > 0 {
		role_id_arr := []int{}
		for _, role_user_vo := range role_user_list {
			role_id_arr = append(role_id_arr, int(role_user_vo.RoleId))
		}
		roleList := []model.Role{}
		service.InitDB().Where(map[string]interface{}{"id": role_id_arr}).Find(&roleList)
		for _, roleVo := range roleList {
			role_name = append(role_name, roleVo.Name)
		}
	}
	ViewData["role_name"] = strings.Join(role_name, ",")
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/personal.html", ViewData)
}

func (uc *UserController) PostDopersonal(ctx *gin.Context) {
	userM := model.User{}
	userM.Nickname = ctx.PostForm("nickname")
	sex, _ := strconv.Atoi(ctx.PostForm("sex"))
	userM.Sex = uint(sex)
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	userM.Age = uint(age)
	userM.Logo = ctx.PostForm("logo")
	userM.Email = ctx.PostForm("email")
	userM.Phone = ctx.PostForm("phone")
	res := service.InitDB().Where(map[string]interface{}{"id": AuthId}).Updates(&userM)
	if res.Error != nil {
		uc.TopjuiError(ctx, "设置失败："+res.Error.Error())
	} else {
		uc.TopjuiSucess(ctx, "设置成功")
	}
}

func (uc *UserController) GetModifypassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/modifypassword.html", ViewData)
}

func (uc *UserController) PostDomodifypassword(ctx *gin.Context) {
	originalPassword := ctx.PostForm("originalPassword")
	password := ctx.PostForm("password")
	password2 := ctx.PostForm("password2")
	if password != password2 {
		uc.TopjuiError(ctx, "新密码两次不一致")
	}
	userM := model.User{}
	service.InitDB().Where(map[string]interface{}{"id": AuthId}).Find(&userM)
	if userM.Password != fmt.Sprintf("%x", md5.Sum([]byte(originalPassword))) {
		uc.TopjuiError(ctx, "原密码不正确,请联系系统管理员重置")
	}
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	res := service.InitDB().Model(userM).Where(map[string]interface{}{"id": AuthId}).Update("password", password)
	if res.Error != nil {
		uc.TopjuiError(ctx, "设置失败："+res.Error.Error())
	} else {
		uc.TopjuiSucess(ctx, "设置成功")
	}
}
