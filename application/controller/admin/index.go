package admin

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/app_session"
	"gocms-gin/config"
	"net/http"
)

type IndexController struct {
	BaseController
}

/**
后台首页
*/
func (c *IndexController) GetIndex(ctx *gin.Context) {
	ViewData["topmenujson"] = c.TopMenu()
	ViewData["authId"] = AuthId
	ViewData["authNickName"] = AuthNickname
	conf := config.InitConfig()
	ViewData["app_icp"] = conf.SiteIcpNum
	ViewData["app_name"] = conf.AppName
	ViewData["site_keywords"] = conf.SiteKeywords
	ViewData["site_description"] = conf.SiteDescription
	ViewData["app_version"] = conf.AppVersion

	ctx.HTML(http.StatusOK, "admin/index/index.html", ViewData)
}

/**
后台主标签页
*/
func (c *IndexController) GetHome(ctx *gin.Context) {
	ViewData["message_count"] = "123,456"
	ViewData["message_today_count"] = "12,345"
	ViewData["pv_count"] = "12,345"
	ViewData["pv_day_count"] = "3,456,789"
	ViewData["user_count"] = "1,000"
	ViewData["user_day_count"] = "1,234"
	ViewData["article_count"] = "456,789"
	ViewData["article_day_count"] = "9,876"

	new_article_list := []map[string]interface{}{
		{"id": 1, "title": "最新文章1"},
		{"id": 2, "title": "最新文章2"},
		{"id": 3, "title": "最新文章3"},
		{"id": 4, "title": "最新文章4"},
		{"id": 5, "title": "最新文章5"},
	}
	hot_article_list := []map[string]interface{}{
		{"id": 1, "title": "热门文章1"},
		{"id": 2, "title": "热门文章2"},
		{"id": 3, "title": "热门文章3"},
		{"id": 4, "title": "热门文章4"},
		{"id": 5, "title": "热门文章5"},
	}
	notice_article_list := []map[string]interface{}{
		{"id": 1, "title": "通知公告1"},
		{"id": 2, "title": "通知公告2"},
		{"id": 3, "title": "通知公告3"},
		{"id": 4, "title": "通知公告4"},
		{"id": 5, "title": "通知公告5"},
	}

	ViewData["new_article_list"] = new_article_list
	ViewData["hot_article_list"] = hot_article_list
	ViewData["notice_article_list"] = notice_article_list
	ctx.HTML(http.StatusOK, "admin/index/home.html", ViewData)
}

func (c *IndexController) GetLogout(ctx *gin.Context) {
	app_session.Sess.Delete("authId")
	app_session.Sess.Clear()
	app_session.Sess.Save()
	//ctx.Redirect(301, "/"+RequestApp+"/public/login")    //301 Chrome会缓存而不访问
	ctx.Redirect(302, "/"+RequestApp+"/public/login")
}
