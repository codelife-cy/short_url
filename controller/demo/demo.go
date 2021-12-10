package demo

import (
	"gin-template/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(group *gin.RouterGroup) {

	group.GET("/test", func(context *gin.Context) {
		utilGin := response.Gin{Ctx: context}
		utilGin.Response(http.StatusOK, "success", "hello world")
	})

}
