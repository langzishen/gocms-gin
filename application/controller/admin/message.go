package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MessageController struct {
	BaseController
}

func (mc *MessageController) GetIndex(ctx *gin.Context) {

}

/**
未读消息个数
*/
func (mc *MessageController) PostNoreadnum(ctx *gin.Context) {
	ctx.Writer.WriteString("0")
}

/**
消息列表
*/
func (mc *MessageController) GetList(ctx *gin.Context) {
	var list []map[string]interface{}
	list = append(list, map[string]interface{}{"id": 1, "text": "..."})
	ctx.JSON(http.StatusOK, map[string]interface{}{"num": 0, "list": list})
}

/**
读消息详情
*/
func (mc *MessageController) PostInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	text := "..."
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id, "text": text})
}
func (mc *MessageController) GeTInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	text := "..."
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id, "text": text})
}
