package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms-gin/application/app_session"
	"gocms-gin/application/extend/admin/rbac"
	"gocms-gin/application/extend/captcha"
	"gocms-gin/application/model"
	"gocms-gin/application/service"
	"gocms-gin/config"
	"net/http"
	"strings"
)

type PublicController struct {
	BaseController
}

/**
登录试图
*/
func (c *PublicController) GetLogin(ctx *gin.Context) {
	app_conf := config.InitConfig()
	if AuthId, _ = app_session.GetInt(app_conf.Rbac.UserAuthKey); AuthId != -1 { //有session中的authId跳转到登录页面
		ctx.Redirect(302, "/"+RequestApp+"/index/index")
	}
	ctx.HTML(http.StatusOK, "admin/public/login.html", ViewData)
}

func (c *PublicController) GetCaptcha(ctx *gin.Context) {
	code := app_session.GetString("admin_login_captcha")
	if code != "" {
		app_session.Sess.Delete("admin_login_captcha")
		app_session.Sess.Save()
	}
	code = captcha.NewCaptcha(ctx.Writer, 120, 40)
	app_session.Set("admin_login_captcha", code)
}

/**
登录操作
*/
func (c *PublicController) PostLoginin(ctx *gin.Context) {
	if is_ajax := ctx.GetHeader("X-Requested-With"); is_ajax == "XMLHttpRequest" {
		data := make(map[string]interface{})
		ctx.BindJSON(&data)
		if "" == data["account"] {
			ctx.JSON(403, map[string]interface{}{"code": 403, "msg": "账号必填!", "data": "", "referer": "/" + RequestApp + "/index/index"})
		}
		if "" == data["password"] {
			ctx.JSON(403, map[string]interface{}{"code": 403, "msg": "账号必填!", "data": "", "referer": "/" + RequestApp + "/index/index"})
		}
		if "" == data["captcha"].(string) {
			ctx.JSON(403, map[string]interface{}{"code": 403, "msg": "请输入验证码!", "data": "", "referer": "/" + RequestApp + "/index/index"})
		}
		code := app_session.GetString("admin_login_captcha")
		if code != "" {
			if strings.ToUpper(code) != strings.ToUpper(data["captcha"].(string)) {
				//ctx.JSON(map[string]interface{}{"code": 403, "msg": "验证码不正确!", "data": "", "referer": "/" + RequestApp + "/public/login"})
				//ctx.StopExecution()
				//测试时不做验证码验证
				//return
			}
		}

		where_map := make(map[string]interface{})
		where_map["account"] = data["account"]
		if isOk, userInfo := new(rbac.RBAC).Authenticate(where_map); !isOk {
			ctx.JSON(403, map[string]interface{}{"code": 403, "msg": "帐号不存在或已禁用！", "data": "", "referer": ""})
		} else {
			if fmt.Sprintf("%x", md5.Sum([]byte(data["password"].(string)))) != userInfo.Password {
				ctx.JSON(403, map[string]interface{}{"code": 403, "msg": "用户名和密码不正确！", "data": "", "referer": ""})
			} else {
				c.saveLoginSession(userInfo)       // 设置后台的SESSION
				c.setLoginLog("boss", userInfo.Id) // 保存登录信息
				//new(rbac.RBAC).SaveAccessList(userInfo.Id) // 缓存访问权限
				ctx.JSON(200, map[string]interface{}{"code": 200, "msg": "", "data": "", "referer": "/" + RequestApp + "/index/index"})
			}
		}
	} else {
		ctx.Redirect(302, "/"+RequestApp+"/index/index")
	}
}

/**
 * 设置后台的SESSION
 */
func (c *PublicController) saveLoginSession(userInfo model.User) {
	appconf := config.InitConfig()
	app_session.Set(appconf.Rbac.UserAuthKey, int(userInfo.Id))
	app_session.Set("email", userInfo.Email)
	app_session.Set("loginUserName", userInfo.Nickname)
	app_session.Set("type_id", userInfo.TypeId)
	if userInfo.TypeId == 9 { //9是超级管理员
		app_session.Set(appconf.Rbac.AdminAuthKey, true)
	} else {
		new(rbac.RBAC).SaveAccessList(userInfo.Id) //设置RBAC权限
	}
}

/**
 * 保存后台的登录信息
 */
func (c *PublicController) setLoginLog(appName string, authId uint) {
	new(service.LogLogin).AddLogLogin(appName, authId)
}
