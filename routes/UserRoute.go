package routes

import (
	"GiveBook/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	//加载静态页面
	//r.LoadHTMLGlob("templates/*")
	//加载资源文件
	//r.Static("/static", "./static")
	r.POST("/api/auth/register", controller.Register)
	return r
}
