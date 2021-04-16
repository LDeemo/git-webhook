package main

import (
	"git-webhook/app"
	"git-webhook/common"
	"github.com/gin-gonic/gin"
)

func init() {
	common.InitLogger()
}

func main() {
	//创建路由
	r := gin.Default()
	//r.Use(common.GinRecovery(common.GetLogger(),true))
	group := r.Group("/hook")
	{
		group.POST("/callback", app.Callback)
	}

	//监听
	r.Run(":9191")
}
