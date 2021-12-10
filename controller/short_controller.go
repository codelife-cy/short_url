package controller

import (
	"gin-template/short"
	"gin-template/util/bind"
	"gin-template/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterShort(group *gin.RouterGroup) {

	/**
	根据长域名获取短域名接口
	*/
	group.POST("/short", func(context *gin.Context) {
		req := new(short.Req)
		bind.Bind(req, context)
		shorter := short.NewShorter()
		url, err := shorter.ShortURL(req.LongURL, req.Note)
		utilGin := response.Gin{Ctx: context}
		if err != nil {
			utilGin.Response(http.StatusBadRequest, "requested url is malformed", nil)
			return
		} else {
			resp := short.Resp{ShortURL: url}
			utilGin.Response(http.StatusOK, "success", resp)
		}
	})

	group.GET("/", func(context *gin.Context) {

	})

}
