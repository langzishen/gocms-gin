package app_session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gocms-gin/config"
	"strconv"
)

var Sess sessions.Session

func SessionAdminStart(group *gin.RouterGroup) {
	var store cookie.Store
	conf := config.InitConfig()
	if conf.SessionRedis {
		store, _ = redis.NewStore(conf.Redis.MaxLineCount, "tcp", conf.Redis.Host+":"+strconv.Itoa(conf.Redis.Port), conf.Redis.Passwrd, []byte("secret_gocms_admin"))
	} else {
		// 创建基于 cookie 的存储引擎，secret_gocms 参数是用于加密的密钥
		store = cookie.NewStore([]byte("secret_gocms_admin"))
	}

	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	group.Use(sessions.Sessions("sessionId", store))
	group.Use(SessStart)
}

func SessStart(ctx *gin.Context) {
	Sess = sessions.Default(ctx)
}

func Set(key, val interface{}) {
	Sess.Set(key, val)
	Sess.Save()
}

type sessE struct {
}

func (e *sessE) Error() string {
	return "no session with key"
}

func GetInt(key interface{}) (val int, err error) {
	res := Sess.Get(key)
	if res == nil {
		return -1, new(sessE)
	} else {
		switch res.(type) {
		case uint:
			val = int(res.(uint))
		case int:
			val = res.(int)
		default:
			val = res.(int)
		}
		err = nil
		return
	}
}

func GetString(key interface{}) (val string) {
	res := Sess.Get(key)
	if res == nil {
		return ""
	} else {
		val = res.(string)
		return
	}
}

func GetBoolean(key string) (val bool, err error) {
	res := Sess.Get(key)
	if res == nil {
		return false, new(sessE)
	} else {
		val = res.(bool)
		err = nil
		return
	}
}
