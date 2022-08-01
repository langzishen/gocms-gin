package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gocms-gin/config"
	"os"
	"strings"
)

type SystemController struct {
	BaseController
}

func (sc *SystemController) GetIndex(ctx *gin.Context) {
	config := config.Config_json
	conf := make(map[string]interface{})
	json.Unmarshal([]byte(config), &conf)
	ViewData["info"] = conf
	ViewData["db"] = conf["db"]
	ViewData["upload"] = conf["upload"]
	ViewData["rbac"] = conf["rbac"]
	sc.BaseController.GetIndex(ctx)
}

func (sc *SystemController) PostIndex(ctx *gin.Context) {
	app_protocol := ctx.PostForm("app_protocol")
	app_url_port := ctx.PostForm("app_url_port")
	app_name := ctx.PostForm("app_name")
	app_url := ctx.PostForm("app_url")
	site_keywords := ctx.PostForm("site_keywords")
	site_description := ctx.PostForm("site_description")
	offline_message := ctx.PostForm("offline_message")
	site_icp_num := ctx.PostForm("site_icp_num")
	app_author_email := ctx.PostForm("app_author_email")
	app_version := ctx.PostForm("app_version")

	db_host := ctx.PostForm("db_host")
	db_port := ctx.PostForm("db_port")
	db_database := ctx.PostForm("db_database")
	db_user := ctx.PostForm("db_user")
	db_user_password := ctx.PostForm("db_user_password")

	file_upload_type := ctx.PostForm("file_upload_type")
	upload_max_size := ctx.PostForm("upload_max_size")
	upload_root_path := ctx.PostForm("upload_root_path")
	upload_region_id := ctx.PostForm("upload_region_id")
	upload_access_key_id := ctx.PostForm("upload_access_key_id")
	upload_access_key_secret := ctx.PostForm("upload_access_key_secret")
	upload_bucket := ctx.PostForm("upload_bucket")
	upload_endpoint := ctx.PostForm("upload_endpoint")
	upload_request_uri := ctx.PostForm("upload_request_uri")
	upload_request_url := ctx.PostForm("upload_request_url")

	rbac_user_auth_type := ctx.PostForm("rbac_user_auth_type")
	rbac_user_auth_key := ctx.PostForm("rbac_user_auth_key")
	rbac_admin_auth_key := ctx.PostForm("rbac_admin_auth_key")
	rbac_require_auth_controller := ctx.PostForm("rbac_require_auth_controller")
	rbac_no_auth_controller := ctx.PostForm("rbac_no_auth_controller")
	rbac_require_auth_action := ctx.PostForm("rbac_require_auth_action")
	rbac_no_auth_action := ctx.PostForm("rbac_no_auth_action")
	rbac_table_access := ctx.PostForm("rbac_table_access")
	rbac_table_user := ctx.PostForm("rbac_table_user")
	rbac_table_node := ctx.PostForm("rbac_table_node")
	rbac_table_role := ctx.PostForm("rbac_table_role")
	rbac_table_role_user := ctx.PostForm("rbac_table_role_user")

	config_demo := config.Config_demo_json
	config_demo = strings.Replace(config_demo, "#$#app_protocol#$#", app_protocol, -1)
	config_demo = strings.Replace(config_demo, "#$#app_url_port#$#", app_url_port, -1)
	config_demo = strings.Replace(config_demo, "#$#app_name#$#", app_name, -1)
	config_demo = strings.Replace(config_demo, "#$#app_url#$#", app_url, -1)
	config_demo = strings.Replace(config_demo, "#$#site_keywords#$#", site_keywords, -1)
	config_demo = strings.Replace(config_demo, "#$#site_description#$#", site_description, -1)
	config_demo = strings.Replace(config_demo, "#$#offline_message#$#", offline_message, -1)
	config_demo = strings.Replace(config_demo, "#$#site_icp_num#$#", site_icp_num, -1)
	config_demo = strings.Replace(config_demo, "#$#app_author_email#$#", app_author_email, -1)
	config_demo = strings.Replace(config_demo, "#$#app_version#$#", app_version, -1)

	config_demo = strings.Replace(config_demo, "#$#host#$#", db_host, -1)
	config_demo = strings.Replace(config_demo, "#$#port#$#", db_port, -1)
	config_demo = strings.Replace(config_demo, "#$#database#$#", db_database, -1)
	config_demo = strings.Replace(config_demo, "#$#user#$#", db_user, -1)
	config_demo = strings.Replace(config_demo, "#$#password#$#", db_user_password, -1)

	config_demo = strings.Replace(config_demo, "#$#file_upload_type#$#", file_upload_type, -1)
	config_demo = strings.Replace(config_demo, "#$#max_size#$#", upload_max_size, -1)
	config_demo = strings.Replace(config_demo, "#$#root_path#$#", upload_root_path, -1)
	config_demo = strings.Replace(config_demo, "#$#region_id#$#", upload_region_id, -1)
	config_demo = strings.Replace(config_demo, "#$#access_key_id#$#", upload_access_key_id, -1)
	config_demo = strings.Replace(config_demo, "#$#access_key_secret#$#", upload_access_key_secret, -1)
	config_demo = strings.Replace(config_demo, "#$#bucket#$#", upload_bucket, -1)
	config_demo = strings.Replace(config_demo, "#$#endpoint#$#", upload_endpoint, -1)
	config_demo = strings.Replace(config_demo, "#$#request_uri#$#", upload_request_uri, -1)
	config_demo = strings.Replace(config_demo, "#$#request_url#$#", upload_request_url, -1)

	config_demo = strings.Replace(config_demo, "#$#user_auth_type#$#", rbac_user_auth_type, -1)
	config_demo = strings.Replace(config_demo, "#$#user_auth_key#$#", rbac_user_auth_key, -1)
	config_demo = strings.Replace(config_demo, "#$#admin_auth_key#$#", rbac_admin_auth_key, -1)
	config_demo = strings.Replace(config_demo, "#$#require_auth_controller#$#", rbac_require_auth_controller, -1)
	config_demo = strings.Replace(config_demo, "#$#no_auth_controller#$#", rbac_no_auth_controller, -1)
	config_demo = strings.Replace(config_demo, "#$#require_auth_action#$#", rbac_require_auth_action, -1)
	config_demo = strings.Replace(config_demo, "#$#no_auth_action#$#", rbac_no_auth_action, -1)
	config_demo = strings.Replace(config_demo, "#$#table_access#$#", rbac_table_access, -1)
	config_demo = strings.Replace(config_demo, "#$#table_user#$#", rbac_table_user, -1)
	config_demo = strings.Replace(config_demo, "#$#table_node#$#", rbac_table_node, -1)
	config_demo = strings.Replace(config_demo, "#$#table_role#$#", rbac_table_role, -1)
	config_demo = strings.Replace(config_demo, "#$#table_role_user#$#", rbac_table_role_user, -1)

	file, err := os.OpenFile("./config/config.json", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		sc.TopjuiError(ctx, "配置文件打开失败")
	}

	_, err = file.WriteString(config_demo)
	if err != nil {
		sc.TopjuiError(ctx, "文件写入失败:"+err.Error())
	}
	sc.TopjuiSucess(ctx, "修改成功")
}
