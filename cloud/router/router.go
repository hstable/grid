package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/controllers/baseStation"
	"github.com/mzz2017/grip/cloud/controllers/company"
	"github.com/mzz2017/grip/cloud/controllers/device"
	"github.com/mzz2017/grip/cloud/controllers/image"
	"github.com/mzz2017/grip/cloud/controllers/line"
	"github.com/mzz2017/grip/cloud/controllers/power"
	"github.com/mzz2017/grip/cloud/controllers/tower"
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
		freeGroup.GET("images/:filename", image.Get)
	}
	adminGroup := Router.Group("api")
	adminGroup.Use(common.JWTAuth(true))
	{
		//只有管理员可以使用的接口
		adminGroup.GET("tree/lines", tree.GetLines)
		adminGroup.GET("tree/line/:id/towers", tree.GetTowers)
		adminGroup.GET("tree/tower/:id/images", tree.GetImages)

		adminGroup.GET("users", user.GetUsers)
		adminGroup.POST("user", user.PostUser)
		adminGroup.PUT("user/:id", user.PutUser)
		adminGroup.DELETE("user/:id", user.DeleteUser)

		adminGroup.GET("companies", company.GetCompanies)
		adminGroup.POST("company", company.Post)
		adminGroup.PUT("company/:id", company.Put)
		adminGroup.DELETE("company/:id", company.Delete)

		adminGroup.GET("powers", power.GetPowers)
		adminGroup.POST("power", power.Post)
		adminGroup.PUT("power/:id", power.Put)
		adminGroup.DELETE("power/:id", power.Delete)

		adminGroup.GET("baseStations", baseStation.GetBaseStations)
		adminGroup.POST("baseStation", baseStation.Post)
		adminGroup.PUT("baseStation/:id", baseStation.Put)
		adminGroup.PUT("baseStation/:id/online", baseStation.PutOnline)
		adminGroup.DELETE("baseStation/:id", baseStation.Delete)

		adminGroup.GET("lines", line.GetLines)
		adminGroup.POST("line", line.Post)
		adminGroup.PUT("line/:id", line.Put)
		adminGroup.DELETE("line/:id", line.Delete)

		adminGroup.GET("line/:lineID/towers", tower.GetTowers)
		adminGroup.POST("line/:lineID/tower", tower.Post)
		adminGroup.PUT("tower/:id", tower.Put)
		adminGroup.DELETE("tower/:id", tower.Delete)

		adminGroup.GET("tower/:towerID/devices", device.GetDevices)
		adminGroup.POST("tower/:towerID/device", device.Post)
		adminGroup.PUT("device/:id", device.Put)
		adminGroup.PUT("device/:id/disabled", device.PutDisabled)
		adminGroup.DELETE("device/:id", device.Delete)

	}
	ordinaryGroup := Router.Group("api")
	ordinaryGroup.Use(common.JWTAuth(false))
	{
		ordinaryGroup.POST("image", image.PostUpload)
	}
}
