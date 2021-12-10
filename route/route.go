package route

import "github.com/gin-gonic/gin"

// Options 定义一个方法类型
type Options func(group *gin.RouterGroup)

var option []Options

func Include(opt ...Options) {
	option = append(opt)
}

func Init() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("")
	for _, opt := range option {
		opt(group)
	}
	return engine
}
