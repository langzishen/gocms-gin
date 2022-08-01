package admin

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/app_session"
	"gocms-gin/application/extend/admin/rbac"
	"gocms-gin/application/extend/upload"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"gocms-gin/config"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var (
	RequestApp        string //访问的当前app
	RequestController string //访问的当前控制器
	RequestAction     string //访问的当前操作
	AuthId            int    //登录者Id
	AuthNickname      string //登陆者昵称
	AuthTypeId        int    //登录者类型，9是超级管理员
	IsAdministrator   bool   //是否是超级管理员
	SearchTids        []int  //数据授权的tids
)
var ViewData = make(gin.H)

type BaseController struct {
}

/**
初始化
*/
func initialize(ctx *gin.Context) {
	ViewData["RequestApp"] = RequestApp
	ViewData["RequestController"] = RequestController
	ViewData["RequestAction"] = RequestAction

	Rbac(ctx)
	SearchTids = DataAccess(ctx)
}

/**
权限认证
*/
func Rbac(ctx *gin.Context) {
	conf := config.InitConfig()
	no_auth_c := strings.Split(conf.Rbac.NoAuthController, ",")
	no_auth_a := strings.Split(conf.Rbac.NoAuthAction, ",")
	if !IsInArr(no_auth_c, RequestController) {
		if !IsInArr(no_auth_a, RequestController+"-"+RequestAction) {
			if !new(rbac.RBAC).AccessDecision(RequestApp, RequestController, RequestAction) {
				TopjuiCtxEnd(ctx)
			}
		}
	}
}

/**
增强型RBAC-数据授权
	请求中带`tid`参数时为操作，将判断该tid是否有数据权限，没有tid时，返回tid的[]int 用于list查询
*/
func DataAccess(ctx *gin.Context) (tids []int) {
	conf := config.InitConfig()

	if !IsAdministrator { //不是超级管理员
		if conf.DataAccess.IsOpen == 1 { //开启了数据授权
			data_access_controllers := strings.Split(conf.DataAccess.Controllers, ",")
			if IsInArr(data_access_controllers, RequestController) {

				role_userM := []model.RoleUser{}
				role_id_s := []int{}
				service.InitDB().Where(map[string]interface{}{"user_id": AuthId}).Find(&role_userM)
				if len(role_userM) > 0 {
					for _, role_user_v := range role_userM {
						role_id_s = append(role_id_s, int(role_user_v.RoleId))
					}
				}

				data_access_map := map[string]interface{}{
					"role_id": role_id_s,
					"model":   RequestController,
				}
				if ctx.Query("node_id") == "" {
					data_access_map["node_id"] = ctx.Query("node_id")
				}
				if ctx.PostForm("node_id") == "" {
					data_access_map["node_id"] = ctx.PostForm("node_id")
				}
				data_accessM := []model.DataAccess{}
				service.InitDB().Where(data_access_map).Find(&data_accessM)
				if len(data_accessM) > 0 {
					for _, data_access_v := range data_accessM {
						tids = append(tids, int(data_access_v.Tid))
					}
				}

				request_tid := 0
				if ctx.Query("tid") != "" {
					request_tid, _ = strconv.Atoi(ctx.Query("tid"))
				}
				if ctx.PostForm("tid") != "" {
					request_tid, _ = strconv.Atoi(ctx.PostForm("tid"))
				}
				if request_tid != 0 {
					if !IntIsInArr(tids, request_tid) {
						TopjuiCtxEnd(ctx)
					}
				}
				return tids
			}
		}
	}
	return []int{}
}

func TopjuiCtxEnd(ctx *gin.Context) {
	if is_ajax := ctx.GetHeader("X-Requested-With"); is_ajax == "XMLHttpRequest" {
		ctx.Header("Content-Type", "application/json; charset=utf-8") //输出JSON头信息
		ctx.JSON(500, map[string]interface{}{"statusCode": 500, "status": 0, "message": "权限不足!", "closeCurrent": true})
	} else {
		//ctx.View("public/no_access.html")
		body := `<div class="topjui-fluid">
						<br>
    					<div class="topjui-row" >
       	 					无访问权限
    					</div>
					</div>`
		ctx.Writer.WriteString(body)
	}
}

/**
输出头部大菜单
*/
func (c *BaseController) TopMenu() []map[string]interface{} {
	appConf := config.InitConfig()
	menu_map_content := map[string]interface{}{"iconCls": appConf.Admin.Menu.Content.Icons, "text": appConf.Admin.Menu.Content.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "content", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	menu_map_framework := map[string]interface{}{"iconCls": appConf.Admin.Menu.Framework.Icons, "text": appConf.Admin.Menu.Framework.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "framework", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	menu_map_groupuser := map[string]interface{}{"iconCls": appConf.Admin.Menu.Groupuser.Icons, "text": appConf.Admin.Menu.Groupuser.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "groupuser", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	menu_map_developers := map[string]interface{}{"iconCls": appConf.Admin.Menu.Developers.Icons, "text": appConf.Admin.Menu.Developers.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "developers", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	menu_map_system := map[string]interface{}{"iconCls": appConf.Admin.Menu.System.Icons, "text": appConf.Admin.Menu.System.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "system", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	menu_map_plugin := map[string]interface{}{"iconCls": appConf.Admin.Menu.Plugin.Icons, "text": appConf.Admin.Menu.Plugin.Title, "codeSetId": "menu", "resourceType": "menu", "levelId": 1, "id": "plugin", "pid": 0, "state": "closed", "url": "", "textColour": ""}
	var top_menu []map[string]interface{}
	top_menu = append(top_menu, menu_map_content, menu_map_framework, menu_map_groupuser, menu_map_developers, menu_map_system, menu_map_plugin)
	return top_menu
}

/**
输出左侧某大菜单下的全部控制器和操作
*/
func (c *BaseController) LeftMenu(ctx *gin.Context) []map[string]interface{} {
	top_menu_item := ctx.Query("menu_id")
	node_list := []*model.Node{}
	var left_menu []map[string]interface{}
	service.InitDB().Where("level=2 AND status=1 and group_id=?", top_menu_item).Order("sort").Find(&node_list)
	for _, node_vo := range node_list {
		node_list2 := []*model.Node{}
		service.InitDB().Where("level=3 AND status=1 AND left_menu_action=1 and pid=?", node_vo.Id).Order("sort").Find(&node_list2)
		i := 0
		for _, node_vo2 := range node_list2 {
			if new(rbac.RBAC).AccessDecision(RequestApp, node_vo.Name, node_vo2.Name) {
				i++
				left_menu = append(left_menu, map[string]interface{}{"iconCls": node_vo2.Icon, "text": node_vo2.Title, "codeSetId": top_menu_item, "resourceType": "menu", "levelId": 3, "id": int(node_vo2.Id), "pid": int(node_vo2.Pid), "state": "open", "url": "/" + RequestApp + "/" + node_vo.Name + "/" + node_vo2.Name, "textColour": ""})
			}
		}
		if i > 0 { //如果控制器下所有的操作都没有权限就不显示这个控制器
			left_menu = append(left_menu, map[string]interface{}{"iconCls": node_vo.Icon, "text": node_vo.Title, "codeSetId": top_menu_item, "resourceType": "menu", "levelId": 2, "id": int(node_vo.Id), "pid": int(node_vo.Pid), "state": "closed", "url": "", "textColour": ""})
		}
	}
	return left_menu
}

/**
输出左侧单个大菜单下的菜单
*/
func (c *BaseController) GetMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.LeftMenu(ctx))
}

/**
全局获取列表页的视图
*/
func (c *BaseController) GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/index.html", ViewData)
}

/**
搜索左侧菜单栏
*/
func (c *BaseController) PostSearchleftmenu(ctx *gin.Context) {
	search := ctx.PostForm("text")
	if search == "" {
		return
	}
	node_list := []*model.Node{}
	var left_menu []map[string]interface{}
	service.InitDB().Where("level=2 AND status=1").Order("sort").Find(&node_list)
	for _, node_vo := range node_list {
		node_list2 := []*model.Node{}
		service.InitDB().Where("level=3 AND status=1 AND left_menu_action=1 and pid=? AND title like '%"+search+"%'", node_vo.Id).Order("sort").Find(&node_list2)
		for _, node_vo2 := range node_list2 {
			if new(rbac.RBAC).AccessDecision(RequestApp, node_vo.Name, node_vo2.Name) {
				left_menu = append(left_menu, map[string]interface{}{"text": node_vo2.Title, "url": "/" + RequestApp + "/" + node_vo.Name + "/" + node_vo2.Name})
			}
		}
	}
	ctx.JSON(http.StatusOK, left_menu)
}

/**
全局添加视图
*/
func (c *BaseController) GetAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/"+"/"+RequestController+"/add.html", ViewData)
}

/**
全局修改视图
*/
func (c *BaseController) GetEdit(ctx *gin.Context) {
	id_s := ctx.Query("id")
	if id_s == "" {
		c.TopjuiError(ctx, "参数Id丢失")
	}
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/edit.html", ViewData)
}

/**
通用删除byID
*/
func (c *BaseController) PostDel(ctx *gin.Context) {
	c.AjaxDone(ctx, "del")
}

func (c *BaseController) PostResume(ctx *gin.Context) {
	c.AjaxDone(ctx, "resume")
}

func (c *BaseController) PostForbid(ctx *gin.Context) {
	c.AjaxDone(ctx, "forbid")
}

func (c *BaseController) PostRecycle(ctx *gin.Context) {
	c.AjaxDone(ctx, "recycle")
}

/**
彻底删除
*/
func (c *BaseController) PostForever_del(ctx *gin.Context) {
	id := ctx.PostForm("id")
	if id == "" {
		c.TopjuiError(ctx, "删除操作参数Id必须")
	}
	var id_arr []int
	for _, id_str := range strings.Split(id, ",") {
		id_i, _ := strconv.Atoi(id_str)
		id_arr = append(id_arr, id_i)
	}
	res := service.InitDB().Exec("DELETE FROM `"+RequestController+"` WHERE id IN ?", id_arr)
	if res.Error == nil {
		c.TopjuiSucess(ctx, "删除成功")
	} else {
		c.TopjuiError(ctx, "删除失败")
	}
}

/**
审核、禁用、删除、恢复
*/
func (c *BaseController) AjaxDone(ctx *gin.Context, do string) {
	id := ctx.PostForm("id")
	if id == "" {
		c.TopjuiError(ctx, "删除操作参数Id必须")
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

	res := service.InitDB().Exec("UPDATE `"+RequestController+"` SET status="+strconv.Itoa(status)+" WHERE id IN ?", id_arr)
	if res.Error == nil {
		c.TopjuiSucess(ctx, say+"成功")
	} else {
		c.TopjuiError(ctx, say+"失败")
	}
}

func (c *BaseController) GetMenu_zree(ctx *gin.Context) {
	node_list := []*model.Node{}
	service.InitDB().Where("level=3 AND status=1 and pid=?", ctx.Query("pid")).Order("sort").Find(&node_list)
	nodeZree_list := []map[string]interface{}{}
	for _, v := range node_list {
		node := model.Node{}
		service.InitDB().Where("id=?", v.Pid).First(&node)
		nodeZree_list = append(nodeZree_list, map[string]interface{}{
			"id":      v.Id,
			"pid":     v.Pid,
			"text":    v.Title,
			"state":   "close",
			"iconCls": "",
			"url":     "/" + RequestApp + "/" + strings.ToLower(node.Name) + "/" + v.Name})
	}
	ctx.JSON(http.StatusOK, nodeZree_list)
}

/**
通用上传
*/
func (c *BaseController) PostAjax_upload(ctx *gin.Context) {
	var upload_file_name string
	if ctx.PostForm("file") != "" {
		upload_file_name = ctx.PostForm("file")
	} else {
		upload_file_name = "file"
	}
	img, _, err := upload.Upload(ctx, upload_file_name) //topjui的单文件上传的文件名默认为file，与input的name无关
	if err != nil {
		ctx.JSON(500, map[string]interface{}{"statusCode": 500, "title": "操作提示", "message": err, "filePath": ""})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{"statusCode": 200, "title": "操作提示", "message": "", "filePath": img})
	}
}

func (c *BaseController) TopjuiSucess(ctx *gin.Context, message string) {
	var navTabTid, forwardUrl string
	if ctx.Request.Method == "GET" {
		navTabTid = ctx.Query("navTabId")
		forwardUrl = ctx.Query("forwardUrl")
	} else {
		navTabTid = ctx.PostForm("navTabId")
		forwardUrl = ctx.PostForm("forwardUrl")
	}
	c.TopjuiReJSON(ctx, 1, message, navTabTid, forwardUrl, true)
}

func (c *BaseController) TopjuiError(ctx *gin.Context, message string) {
	var navTabTid, forwardUrl string
	if ctx.Request.Method == "GET" {
		navTabTid = ctx.Query("navTabId")
		forwardUrl = ctx.Query("forwardUrl")
	} else {
		navTabTid = ctx.PostForm("navTabId")
		forwardUrl = ctx.PostForm("forwardUrl")
	}
	c.TopjuiReJSON(ctx, 0, message, navTabTid, forwardUrl, false)
}

/**
 *  code返回的状态码
 */
func (c *BaseController) TopjuiReJSON(ctx *gin.Context, code int, message string, navTabTid string, forwardUrl string, callbackType bool) {
	res := make(map[string]interface{})
	res["status"] = code
	if code == 1 {
		res["statusCode"] = 200
	} else {
		if code == 0 {
			res["statusCode"] = 300
		} else {
			res["statusCode"] = code
		}
	}
	res["message"] = message
	res["tabid"] = navTabTid
	res["forwardUrl"] = forwardUrl
	res["closeCurrent"] = callbackType
	ctx.Header("Content-Type", "application/json; charset=utf-8") //输出JSON头信息
	ctx.JSON(http.StatusOK, res)
}

/****************** 通用函数 ***************************/
/**
字符串是否在数组里
*/
func IsInArr(arr []string, str string) (is bool) {
	is = false
	for _, vo := range arr {
		if vo == str {
			is = true
			return is
		}
	}
	return is
}

/**
数字是否在数组里
*/
func IntIsInArr(arr []int, i int) (is bool) {
	is = false
	for _, vo := range arr {
		if vo == i {
			is = true
			return is
		}
	}
	return is
}

/**
去重[]map[string]int
*/
func unique_slice_map(slice_map []map[string]int) []map[string]int {
	s_len := len(slice_map) - 1
	for ; s_len > 0; s_len-- {
		for j := s_len - 1; j >= 0; j-- {
			if cmpMap(slice_map[s_len], slice_map[j]) {
				slice_map = append(slice_map[:s_len], slice_map[s_len+1:]...)
				break
			}
		}
	}
	return slice_map
}

/**
比较两个map[string]int是否相同
*/
func cmpMap(m1, m2 map[string]int) bool {
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, ok := m1[k2]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func RandString(n int) string {
	allstring := "qwertyuiopasdfghjklzxcvbnm1233456789"
	ret := ""
	for i := 0; i < n; i++ {
		tmp := rand.Intn(len(allstring))
		ret += allstring[tmp : tmp+1]
	}
	return ret
}

/***********************通用函数****END**********************/

/**
后台登录权限验证中间件
*/
func AdminAuth(ctx *gin.Context) {
	app_conf := config.InitConfig()
	full_path := ctx.FullPath()
	pathArr := strings.Split(full_path, "/")
	RequestApp = pathArr[1]        //app
	RequestController = pathArr[2] //controller
	RequestAction = pathArr[3]     //action
	if authId, err := app_session.GetInt(app_conf.Rbac.UserAuthKey); err == nil && authId != -1 {
		AuthId = authId
		AuthTypeId, _ = app_session.GetInt("type_id")
		if AuthTypeId == 9 { //超级管理员
			AuthNickname = app_session.GetString("loginUserName") + "(超级管理员)"
			IsAdministrator = true
		} else {
			AuthNickname = app_session.GetString("loginUserName")
			IsAdministrator = false
		}
		initialize(ctx)
	} else { //没有session中的authId
		if (RequestController == "public" && RequestAction == "login") || (RequestController == "public" && RequestAction == "loginin") || (RequestController == "public" && RequestAction == "captcha") { //登录操作不做权限验证
			initialize(ctx)
		} else {
			//跳转到登录
			//ctx.Writef("no login")
			ctx.Redirect(302, "/"+RequestApp+"/public/login")
		}
	}
}
