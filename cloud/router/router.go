package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/controllers/image"
	"github.com/mzz2017/grip/cloud/controllers/line"
	"github.com/mzz2017/grip/cloud/controllers/tree"
	"github.com/mzz2017/grip/cloud/controllers/user"
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
		freeGroup.POST("user/login", user.PostLogin)
		freeGroup.GET("images/:filename", image.GetImage)
	}
	adminGroup := Router.Group("api")
	adminGroup.Use(common.JWTAuth(true))
	{
		//只有管理员可以使用的接口
		adminGroup.POST("user/register", user.PostRegister)
		adminGroup.POST("line", line.PostLine)
		adminGroup.GET("lines", tree.GetLines)
		adminGroup.GET("line/:id/towers", tree.GetTowers)
		adminGroup.GET("tower/:id/images", tree.GetImages)
	}
	ordinaryGroup := Router.Group("api")
	ordinaryGroup.Use(common.JWTAuth(true))
	{
		ordinaryGroup.POST("image", image.PostUpload)
	}
}
