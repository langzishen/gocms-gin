package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"gocms-gin/config"
	"net/http"
	"strconv"
)

type NodeController struct {
	BaseController
}

func (nc *NodeController) PostTop_menu(ctx *gin.Context) {
	top_menu := nc.TopMenu()
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": 1, "total": len(top_menu), "rows": top_menu})
}

func (nc *NodeController) PostController_list(ctx *gin.Context) {
	group_id := ctx.PostForm("group_id")
	if group_id == "" {
		nc.TopjuiError(ctx, "大菜单ID必须")
	}
	var controller_menu []map[string]interface{}
	node_list := []*model.Node{}
	service.InitDB().Where("level=2 AND group_id=?", group_id).Order("sort").Find(&node_list)
	for _, node_vo := range node_list {
		controller_menu = append(controller_menu, map[string]interface{}{"iconCls": "fa fa-" + node_vo.Icon, "name": node_vo.Name, "text": node_vo.Title, "status": node_vo.Status, "codeSetId": group_id, "resourceType": "menu", "levelId": 2, "id": int(node_vo.Id), "pid": int(node_vo.Pid), "state": "closed", "url": "", "textColour": ""})
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": 1, "total": len(controller_menu), "rows": controller_menu})
}

func (nc *NodeController) PostAction_list(ctx *gin.Context) {
	group_id := ctx.PostForm("group_id")
	controller_id, _ := strconv.Atoi(ctx.PostForm("controller_id"))
	if controller_id == 0 {
		nc.TopjuiError(ctx, "控制ID必须")
	}
	var action_menu []map[string]interface{}
	node_list := []*model.Node{}
	service.InitDB().Where("level=3 AND pid=?", controller_id).Order("sort").Find(&node_list)
	for _, node_vo := range node_list {
		action_menu = append(action_menu, map[string]interface{}{"iconCls": "fa fa-" + node_vo.Icon, "name": node_vo.Name, "text": node_vo.Title, "status": node_vo.Status, "codeSetId": group_id, "resourceType": "menu", "levelId": 2, "id": int(node_vo.Id), "pid": int(node_vo.Pid), "state": "closed", "url": "", "textColour": ""})
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"pages": 1, "total": len(action_menu), "rows": action_menu})
}

/**
添加视图
*/
func (nc *NodeController) GetController_add(ctx *gin.Context) {
	group_id := ctx.Query("group_id")
	ViewData["group_id"] = group_id
	appConf := config.InitConfig()
	var group_name string
	switch group_id {
	case "content":
		group_name = appConf.Admin.Menu.Content.Title
	case "framework":
		group_name = appConf.Admin.Menu.Framework.Title
	case "groupuser":
		group_name = appConf.Admin.Menu.Groupuser.Title
	case "developers":
		group_name = appConf.Admin.Menu.Developers.Title
	case "system":
		group_name = appConf.Admin.Menu.System.Title
	case "plugin":
		group_name = appConf.Admin.Menu.Plugin.Title
	default:
		group_name = ""
	}
	ViewData["group_name"] = group_name
	/**
	top_menu := nc.TopMenu()
	var top_menu_list []map[string]string
	for _, top_menu_vo := range top_menu {
		top_menu_list = append(top_menu_list, map[string]string{"text": top_menu_vo["text"].(string), "value": top_menu_vo["id"].(string)})
	}
	ctx.ViewData("top_menu_list", top_menu_list)**/
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/controller_add.html", ViewData)
}

func (nc *NodeController) PostController_add(ctx *gin.Context) {
	post_data := make(map[string]interface{})
	post_data["name"] = ctx.PostForm("name")
	post_data["title"] = ctx.PostForm("title")
	post_data["status"], _ = strconv.Atoi(ctx.PostForm("status"))
	post_data["remark"] = ctx.PostForm("remark")
	post_data["sort"], _ = strconv.Atoi(ctx.PostForm("sort"))
	post_data["pid"], _ = strconv.Atoi(ctx.PostForm("pid"))
	post_data["level"], _ = strconv.Atoi(ctx.PostForm("level"))
	post_data["type"] = 0
	post_data["group_id"] = ctx.PostForm("group_id")
	post_data["icon"] = ctx.PostForm("icon")
	post_data["left_menu_action"] = 0

	post_data_byte, _ := json.Marshal(post_data)
	nodeM := model.Node{}
	json.Unmarshal(post_data_byte, &nodeM)
	service.InitDB().Create(&nodeM)
	if nodeM.Id > 0 {
		nc.TopjuiSucess(ctx, "添加成功")
	} else {
		nc.TopjuiError(ctx, "添加失败")
	}
}

/**
修改视图
*/
func (nc *NodeController) GetController_edit(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		nc.TopjuiError(ctx, "修改参数Id必须")
	}
	nodeM := model.Node{}
	id_i, _ := strconv.Atoi(id)
	service.InitDB().Where(map[string]interface{}{"id": id_i}).Last(&nodeM)
	ViewData["info"] = nodeM
	node_upM := model.Node{}
	service.InitDB().Where(map[string]interface{}{"id": nodeM.Pid}).Last(&node_upM)
	ViewData["up_info"] = node_upM
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/controller_edit.html", ViewData)
}

/**
处理修改
*/
func (nc *NodeController) PostController_edit(ctx *gin.Context) {
	if ctx.PostForm("id") == "" {
		nc.TopjuiError(ctx, "参数Id必须")
	}
	post_data := make(map[string]interface{})
	post_data["id"], _ = strconv.Atoi(ctx.PostForm("id"))
	post_data["name"] = ctx.PostForm("name")
	post_data["title"] = ctx.PostForm("title")
	post_data["status"], _ = strconv.Atoi(ctx.PostForm("status"))
	post_data["remark"] = ctx.PostForm("remark")
	post_data["sort"], _ = strconv.Atoi(ctx.PostForm("sort"))
	post_data["pid"], _ = strconv.Atoi(ctx.PostForm("pid"))
	post_data["level"], _ = strconv.Atoi(ctx.PostForm("level"))
	post_data["type"] = 0
	post_data["group_id"] = ctx.PostForm("group_id")
	post_data["icon"] = ctx.PostForm("icon")
	post_data["left_menu_action"] = 0

	post_data_byte, _ := json.Marshal(post_data)
	nodeM := model.Node{}
	json.Unmarshal(post_data_byte, &nodeM)
	res := service.InitDB().Save(&nodeM)
	if res.Error == nil {
		nc.TopjuiSucess(ctx, "保存成功")
	} else {
		nc.TopjuiError(ctx, "保存失败")
	}
}

func (nc *NodeController) GetAction_add(ctx *gin.Context) {
	p_id := ctx.Query("pid")
	p_id_i, _ := strconv.Atoi(p_id)
	ViewData["p_id"] = p_id_i

	nodeM := model.Node{}
	service.InitDB().Where(map[string]interface{}{"id": p_id_i}).Last(&nodeM)
	ViewData["p_title"] = nodeM.Title

	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/action_add.html", ViewData)
}

func (nc *NodeController) PostAction_add(ctx *gin.Context) {
	post_data := make(map[string]interface{})
	post_data["name"] = ctx.PostForm("name")
	post_data["title"] = ctx.PostForm("title")
	post_data["status"], _ = strconv.Atoi(ctx.PostForm("status"))
	post_data["remark"] = ctx.PostForm("remark")
	post_data["sort"] = 999
	post_data["pid"], _ = strconv.Atoi(ctx.PostForm("pid"))
	post_data["level"], _ = strconv.Atoi(ctx.PostForm("level"))
	post_data["type"] = 0
	post_data["group_id"] = ctx.PostForm("group_id")
	post_data["icon"] = ctx.PostForm("icon")
	post_data["left_menu_action"], _ = strconv.Atoi(ctx.PostForm("left_menu_action"))

	post_data_byte, _ := json.Marshal(post_data)
	nodeM := model.Node{}
	json.Unmarshal(post_data_byte, &nodeM)
	service.InitDB().Create(&nodeM)
	if nodeM.Id > 0 {
		nc.TopjuiSucess(ctx, "添加成功")
	} else {
		nc.TopjuiError(ctx, "添加失败")
	}
}

/**
修改视图
*/
func (nc *NodeController) GetAction_edit(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		nc.TopjuiError(ctx, "修改参数Id必须")
	}
	nodeM := model.Node{}
	id_i, _ := strconv.Atoi(id)
	service.InitDB().Where(map[string]interface{}{"id": id_i}).Last(&nodeM)
	ViewData["info"] = nodeM
	node_upM := model.Node{}
	service.InitDB().Where(map[string]interface{}{"id": nodeM.Pid}).Last(&node_upM)
	ViewData["up_info"] = node_upM
	ctx.HTML(http.StatusOK, "admin/"+RequestController+"/action_edit.html", ViewData)
}

/**
处理修改
*/
func (nc *NodeController) PostAction_edit(ctx *gin.Context) {
	if ctx.PostForm("id") == "" {
		nc.TopjuiError(ctx, "参数Id必须")
	}
	post_data := make(map[string]interface{})
	post_data["id"], _ = strconv.Atoi(ctx.PostForm("id"))
	post_data["name"] = ctx.PostForm("name")
	post_data["title"] = ctx.PostForm("title")
	post_data["status"], _ = strconv.Atoi(ctx.PostForm("status"))
	post_data["remark"] = ctx.PostForm("remark")
	post_data["sort"], _ = strconv.Atoi(ctx.PostForm("sort"))
	post_data["pid"], _ = strconv.Atoi(ctx.PostForm("pid"))
	post_data["level"], _ = strconv.Atoi(ctx.PostForm("level"))
	post_data["type"] = 0
	post_data["group_id"] = ctx.PostForm("group_id")
	post_data["icon"] = ctx.PostForm("icon")
	post_data["left_menu_action"], _ = strconv.Atoi(ctx.PostForm("left_menu_action"))

	post_data_byte, _ := json.Marshal(post_data)
	nodeM := model.Node{}
	json.Unmarshal(post_data_byte, &nodeM)
	res := service.InitDB().Save(&nodeM)
	if res.Error == nil {
		nc.TopjuiSucess(ctx, "保存成功")
	} else {
		nc.TopjuiError(ctx, "保存失败")
	}
}
