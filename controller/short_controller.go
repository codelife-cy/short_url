package controller

import (
	"gin-template/short"
	"gin-template/util/bind"
	"gin-template/util/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
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

	/**
	访问短域名重定向到长域名地址
	*/
	group.GET("/:shortURL", func(context *gin.Context) {
		param := context.Param("shortURL")
		pat := `^[a-zA-Z0-9]{1,11}$`
		compile := regexp.MustCompile(pat)
		// 判断是否是短域名请求
		if ok := compile.MatchString(param); !ok {
			return
		}
		shorter := short.NewShorter()
		expand, err := shorter.Expand(param)
		if err != nil {
			log.Printf("redirect short url error. %v", err)
			context.Error(err)
			return
		}
		if len(expand) != 0 {
			context.Redirect(http.StatusTemporaryRedirect, expand)
		} else {
			context.Writer.WriteHeader(http.StatusNoContent)
		}
	})

}
