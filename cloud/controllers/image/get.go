package image

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/config"
	"path"
)

func Get(ctx *gin.Context) {
	fname := ctx.Param("filename")
	ctx.File(path.Join(config.Get().AssetDir, fname))
}
