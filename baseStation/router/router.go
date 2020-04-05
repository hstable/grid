package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/basestation/common"
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
		userGroup := freeGroup.Group("user")
	}
	Router.GET("images/:hash")
	adminGroup := Router.Group("api")
	adminGroup.Use(common.JWTAuth(true))
	{
		//只有管理员可以使用的接口
		userGroup := adminGroup.Group("user")
		{
			userGroup.POST("register", user.PostRegister)
		}
		imageGroup := adminGroup.Group("image")
		{
			imageGroup.GET("list", image.GetList)
			imageGroup.GET("next", image.GetNext)
			imageGroup.GET("last", image.GetLast)
			imageGroup.PUT("mark", image.PutMark)
		}
	}
}
