package main

import (
	"gin-template/config"
	"gin-template/controller"
	"gin-template/controller/demo"
	"gin-template/route"
)

func main() {
	cfg := config.Get()
	route.Include(demo.Router, controller.RouterShort)
	engine := route.Init()

	engine.Run(cfg.Server.Port)
}
