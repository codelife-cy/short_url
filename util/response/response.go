package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, response{Code: code, Message: msg, Data: data})
	return
}
