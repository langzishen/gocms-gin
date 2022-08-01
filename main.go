package main

import (
	"github.com/gin-gonic/gin"
	"gocms-gin/application/app_session"
	"gocms-gin/application/controller/admin"
	"gocms-gin/application/controller/index"
	"gocms-gin/application/extend/upload"
	"gocms-gin/config"
	"strconv"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")            //静态资源
	router.LoadHTMLGlob("application/view/**/**/*") //视图模板
	upload.MaxMultipartMemory = router              //将router传给上传扩展包用于设置最大上传size

	admin_router := router.Group("boss")
	app_session.SessionAdminStart(admin_router) //开启session
	admin_router.Use(admin.AdminAuth)           //使用验证中间件
	RouterAdmin(admin_router)

	index_router := router.Group("index")
	app_session.SessionAdminStart(index_router) //开启session
	index_router.Use(index.IndexAuth)           //使用验证中间件
	RouterIndex(index_router)

	router.Run(":" + strconv.Itoa(config.InitConfig().AppUrlPort))
}

/**
后台部分的路由
*/
func RouterAdmin(group *gin.RouterGroup) {
	group.GET("/public/login", new(admin.PublicController).GetLogin)
	group.GET("/public/captcha", new(admin.PublicController).GetCaptcha)
	group.POST("/public/loginin", new(admin.PublicController).PostLoginin)

	group.GET("/index/index", new(admin.IndexController).GetIndex)
	group.POST("/index/searchleftmenu", new(admin.IndexController).PostSearchleftmenu)
	group.GET("/index/menu", new(admin.IndexController).GetMenu)
	group.GET("/index/home", new(admin.IndexController).GetHome)
	group.GET("/index/logout", new(admin.IndexController).GetLogout)

	group.GET("/article/index", new(admin.ArticleController).GetIndex)
	group.POST("/article/list", new(admin.ArticleController).PostList)
	group.GET("/article/add", new(admin.ArticleController).GetAdd)
	group.POST("/article/add", new(admin.ArticleController).PostAdd)
	group.GET("/article/edit", new(admin.ArticleController).GetEdit)
	group.POST("/article/edit", new(admin.ArticleController).PostEdit)
	group.POST("/article/del", new(admin.ArticleController).PostDel)
	group.POST("/article/forever_del", new(admin.ArticleController).PostForever_del)
	group.POST("/article/forbid", new(admin.ArticleController).PostForbid)
	group.POST("/article/recycle", new(admin.ArticleController).PostRecycle)
	group.POST("/article/resume", new(admin.ArticleController).PostResume)
	group.POST("/article/ajax_upload", new(admin.ArticleController).PostAjax_upload)
	group.POST("/article/zree", new(admin.ArticleController).PostZree)

	group.GET("/attribute/index", new(admin.AttributeController).GetIndex)
	group.POST("/attribute/list", new(admin.AttributeController).PostList)
	group.GET("/attribute/add", new(admin.AttributeController).GetAdd)
	group.POST("/attribute/add", new(admin.AttributeController).PostAdd)
	group.GET("/attribute/edit", new(admin.AttributeController).GetEdit)
	group.POST("/attribute/edit", new(admin.AttributeController).PostEdit)
	group.POST("/attribute/del", new(admin.AttributeController).PostDel)
	group.POST("/attribute/forever_del", new(admin.AttributeController).PostForever_del)
	group.POST("/attribute/forbid", new(admin.AttributeController).PostForbid)
	group.POST("/attribute/recycle", new(admin.AttributeController).PostRecycle)
	group.POST("/attribute/resume", new(admin.AttributeController).PostResume)

	group.GET("/category/index", new(admin.CategoryController).GetIndex)
	group.POST("/category/list", new(admin.CategoryController).PostList)
	group.GET("/category/add", new(admin.CategoryController).GetAdd)
	group.POST("/category/add", new(admin.CategoryController).PostAdd)
	group.GET("/category/edit", new(admin.CategoryController).GetEdit)
	group.POST("/category/edit", new(admin.CategoryController).PostEdit)
	group.POST("/category/del", new(admin.CategoryController).PostDel)
	group.POST("/category/forever_del", new(admin.CategoryController).PostForever_del)
	group.POST("/category/forbid", new(admin.CategoryController).PostForbid)
	group.POST("/category/recycle", new(admin.CategoryController).PostRecycle)
	group.POST("/category/resume", new(admin.CategoryController).PostResume)

	group.GET("/message/index", new(admin.MessageController).GetIndex)
	group.GET("/message/list", new(admin.MessageController).GetList)
	group.GET("/message/add", new(admin.MessageController).GetAdd)
	group.GET("/message/edit", new(admin.MessageController).GetEdit)
	group.POST("/message/del", new(admin.MessageController).PostDel)
	group.POST("/message/forever_del", new(admin.MessageController).PostForever_del)
	group.POST("/message/forbid", new(admin.MessageController).PostForbid)
	group.POST("/message/recycle", new(admin.MessageController).PostRecycle)
	group.POST("/message/resume", new(admin.MessageController).PostResume)
	group.POST("/message/noreadnum", new(admin.MessageController).PostNoreadnum)
	group.GET("/message/info", new(admin.MessageController).GeTInfo)
	group.POST("/message/info", new(admin.MessageController).PostInfo)

	group.GET("/model/index", new(admin.ModelController).GetIndex)
	group.POST("/model/list", new(admin.ModelController).PostList)
	group.GET("/model/add", new(admin.ModelController).GetAdd)
	group.POST("/model/add", new(admin.ModelController).PostAdd)
	group.GET("/model/edit", new(admin.ModelController).GetEdit)
	group.POST("/model/edit", new(admin.ModelController).PostEdit)
	group.POST("/model/del", new(admin.ModelController).PostDel)
	group.POST("/model/forever_del", new(admin.ModelController).PostForever_del)
	group.POST("/model/forbid", new(admin.ModelController).PostForbid)
	group.POST("/model/recycle", new(admin.ModelController).PostRecycle)
	group.POST("/model/resume", new(admin.ModelController).PostResume)

	group.GET("/node/index", new(admin.NodeController).GetIndex)
	group.POST("/node/top_menu", new(admin.NodeController).PostTop_menu)
	group.POST("/node/controller_list", new(admin.NodeController).PostController_list)
	group.POST("/node/action_list", new(admin.NodeController).PostAction_list)
	group.GET("/node/controller_add", new(admin.NodeController).GetController_add)
	group.POST("/node/controller_add", new(admin.NodeController).PostController_add)
	group.GET("/node/controller_edit", new(admin.NodeController).GetController_edit)
	group.POST("/node/controller_edit", new(admin.NodeController).PostController_edit)
	group.GET("/node/action_add", new(admin.NodeController).GetAction_add)
	group.POST("/node/action_add", new(admin.NodeController).PostAction_add)
	group.GET("/node/action_edit", new(admin.NodeController).GetAction_edit)
	group.POST("/node/action_edit", new(admin.NodeController).PostAction_edit)
	group.POST("/node/del", new(admin.NodeController).PostDel)
	group.POST("/node/forever_del", new(admin.NodeController).PostForever_del)
	group.POST("/node/forbid", new(admin.NodeController).PostForbid)
	group.POST("/node/recycle", new(admin.NodeController).PostRecycle)
	group.POST("/node/resume", new(admin.NodeController).PostResume)

	group.GET("/photo/index", new(admin.PhotoController).GetIndex)
	group.POST("/photo/list", new(admin.PhotoController).PostList)
	group.GET("/photo/add", new(admin.PhotoController).GetAdd)
	group.POST("/photo/add", new(admin.PhotoController).PostAdd)
	group.GET("/photo/edit", new(admin.PhotoController).GetEdit)
	group.POST("/photo/edit", new(admin.PhotoController).PostEdit)
	group.POST("/photo/del", new(admin.PhotoController).PostDel)
	group.POST("/photo/forever_del", new(admin.PhotoController).PostForever_del)
	group.POST("/photo/forbid", new(admin.PhotoController).PostForbid)
	group.POST("/photo/recycle", new(admin.PhotoController).PostRecycle)
	group.POST("/photo/resume", new(admin.PhotoController).PostResume)
	group.POST("/photo/ajax_upload", new(admin.PhotoController).PostAjax_upload)
	group.POST("/photo/zree", new(admin.PhotoController).PostZree)

	group.GET("/role/index", new(admin.RoleController).GetIndex)
	group.POST("/role/list", new(admin.RoleController).PostList)
	group.GET("/role/add", new(admin.RoleController).GetAdd)
	group.POST("/role/add", new(admin.RoleController).PostAdd)
	group.GET("/role/edit", new(admin.RoleController).GetEdit)
	group.POST("/role/edit", new(admin.RoleController).PostEdit)
	group.POST("/role/del", new(admin.RoleController).PostDel)
	group.POST("/role/forever_del", new(admin.RoleController).PostForever_del)
	group.POST("/role/forbid", new(admin.RoleController).PostForbid)
	group.POST("/role/recycle", new(admin.RoleController).PostRecycle)
	group.POST("/role/resume", new(admin.RoleController).PostResume)
	group.POST("/role/ajax_upload", new(admin.RoleController).PostAjax_upload)
	group.POST("/role/zree", new(admin.RoleController).PostZree)
	group.GET("/role/access", new(admin.RoleController).GetAccess)
	group.POST("/role/accesstree", new(admin.RoleController).PostAccesstree)
	group.GET("/role/doaccess", new(admin.RoleController).GetDoaccess)
	group.GET("/role/roleuser", new(admin.RoleController).GetRoleuser)
	group.POST("/role/roleuser", new(admin.RoleController).GetRoleuser)
	group.GET("/role/dataaccess", new(admin.RoleController).GetDataaccess)
	group.POST("/role/data_access_controllers", new(admin.RoleController).PostData_access_controllers)
	group.POST("/role/data_access_actions", new(admin.RoleController).PostData_access_actions)
	group.POST("/role/data_access_tids", new(admin.RoleController).PostData_access_tids)
	group.POST("/role/doadddataaccess", new(admin.RoleController).PostDoadddataaccess)
	group.POST("/role/dodeldataaccess", new(admin.RoleController).PostDodeldataaccess)

	group.GET("/static/index", new(admin.StaticController).GetIndex)
	group.POST("/static/list", new(admin.StaticController).PostList)
	group.GET("/static/add", new(admin.StaticController).GetAdd)
	group.GET("/static/edit", new(admin.StaticController).GetEdit)
	group.POST("/static/del", new(admin.StaticController).PostDel)
	group.POST("/static/forever_del", new(admin.StaticController).PostForever_del)
	group.POST("/static/forbid", new(admin.StaticController).PostForbid)
	group.POST("/static/recycle", new(admin.StaticController).PostRecycle)
	group.POST("/static/resume", new(admin.StaticController).PostResume)
	group.POST("/static/adddir", new(admin.StaticController).PostAdddir)
	group.POST("/static/doadddir", new(admin.StaticController).PostDoadddir)
	group.POST("/static/multiupload", new(admin.StaticController).PostMultiupload)
	group.POST("/static/delfile", new(admin.StaticController).PostDelfile)
	group.GET("/static/download", new(admin.StaticController).GetDownload)

	group.GET("/system/index", new(admin.SystemController).GetIndex)
	group.POST("/system/index", new(admin.SystemController).PostIndex)

	group.GET("/user/index", new(admin.UserController).GetIndex)
	group.POST("/user/list", new(admin.UserController).PostList)
	group.GET("/user/add", new(admin.UserController).GetAdd)
	group.POST("/user/add", new(admin.UserController).PostAdd)
	group.GET("/user/edit", new(admin.UserController).GetEdit)
	group.POST("/user/edit", new(admin.UserController).PostEdit)
	group.POST("/user/del", new(admin.UserController).PostDel)
	group.POST("/user/forever_del", new(admin.UserController).PostForever_del)
	group.POST("/user/forbid", new(admin.UserController).PostForbid)
	group.POST("/user/recycle", new(admin.UserController).PostRecycle)
	group.POST("/user/resume", new(admin.UserController).PostResume)
	group.POST("/user/ajax_upload", new(admin.UserController).PostAjax_upload)
	group.GET("/user/resetpassword", new(admin.UserController).GetResetpassword)
	group.POST("/user/resetpassword", new(admin.UserController).PostResetpassword)
	group.GET("/user/personal", new(admin.UserController).GetPersonal)
	group.POST("/user/dopersonal", new(admin.UserController).PostDopersonal)
	group.GET("/user/modifypassword", new(admin.UserController).GetModifypassword)
	group.POST("/user/domodifypassword", new(admin.UserController).PostDomodifypassword)

}

/**
前台部分的路由
*/
func RouterIndex(group *gin.RouterGroup) {
	group.GET("/index/index", new(index.IndexController).GetIndex)

	group.GET("/about/index", new(index.AboutController).GetIndex)
	group.GET("/about/join", new(index.AboutController).GetJoin)
	group.GET("/about/joininfo", new(index.AboutController).GetJoininfo)

	group.GET("/article/joininfo", new(index.ArticleController).GetIndex)

	group.GET("/contact/joininfo", new(index.ContactController).GetIndex)

	group.GET("/product/joininfo", new(index.ProductController).GetIndex)
}
