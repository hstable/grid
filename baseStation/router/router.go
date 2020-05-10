package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/basestation/controllers"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 跨域
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("DELETE", "PUT")
	Router.Use(cors.New(corsConfig))
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	Router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	Router.Use(gin.Recovery())
	freeGroup := Router.Group("api")
	{
		//无需登录也可使用的接口
		freeGroup.GET("image/:filename", controllers.GetImage)
		freeGroup.POST("devices", controllers.PostDevices)
		freeGroup.GET("device/:ip/newPhoto", controllers.GetNewPhoto)
		freeGroup.GET("log", controllers.GetLog)
	}
}
