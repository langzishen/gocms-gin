# gocms

gocms是go语言实现的一套cms内容管理系统,服务端使用go语言,主体框架使用[gin](https://github.com/gin-gonic/gin)框架,采用mvc架构,管理后台前端使用[topjui](https://www.topjui.com/),数据库使用mysql,ORM使用[goorm](https://gorm.io),内置RBAC、UPLOAD、OSS等扩展。

#### 环境搭建：

	1.go version:1.16+
	2.首次运行需修改项目根目录下config/config.json配置文件。项目运行起来后可以在后台的系统管理中修改部分配置。
	3.导入数据库文件:db/gocms.sql到mysql中,数据库名称需与config.json保持一致
	4.启动方式：go run main.go
	5.前台访问地址：域名/index/index/index
	6.后台访问地址：域名/boss/index/index,(默认)，也可以在main.go中修改`router.Group("boss")`,中的`“boss”`为其他路径。
	7.后台的超级管理员账号/密码：admin/amin
	8.`/static/uploads`需要有写入权限。



#### 主要包依赖：

* github.com/gin-gonic/gin v1.8.1
* github.com/gin-contrib/sessions v0.0.5
* gorm.io/driver/mysql
* github.com/go-sql-driver/mysql
* gorm.io/gorm
* github.com/aliyun/aliyun-oss-go-sdk


## 运行：
* 首次运行前，修改项目config/config.json文件，修改对应的访问域名、端口号、数据库等信息，端口号如果不修改默认为8080，域名如果不设置默认为127.0.0.1，数据库sql文件在db目录下，配置好数据库就可以运行了，首次运行后，访问域名:端口号/boss/index/index，进入后台在系统设置里面可以进行更改编辑操作。
* 在项目的根目录下执行:go run main.go
* 后台访问地址:域名:端口号/boss/index/index(例：127.0.0.1:8080/boss/index/index),如需修改boss路径，请修改main.go文件的19行`admin_router := router.Group("boss")`为`admin_router := router.Group("xxxx")`,xxxx:你想要设置的后台访问路径。
* 前台访问地址：域名:端口号/index/index/index(例：127.0.0.1:8080/index/index/index)，由于本系统主要是后台系统，前台只是测试展示了一下效果，开发者可以根据自己喜好编写。
