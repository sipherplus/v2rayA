package handlers

import (
	"github.com/gin-gonic/gin"
	"v2rayA/tools"
)

func Version(ctx *gin.Context) {
	tools.ResponseSuccess(ctx, "1.00")
}