package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/basestation/config"
	"path"
)

func GetImage(ctx *gin.Context) {
	fname := ctx.Param("filename")
	ctx.File(path.Join(config.Get().AssetDir, fname))
}
