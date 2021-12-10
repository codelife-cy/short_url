package bind

import "github.com/gin-gonic/gin"

// Bind 参数绑定
func Bind(model interface{}, ctx *gin.Context) (interface{}, error) {
	if err := ctx.ShouldBind(model); err != nil {
		return nil, err
	}
	return model, nil
}
