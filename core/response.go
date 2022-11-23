package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseWithSuccess(ctx *gin.Context, message string, result interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"success": true,
		"result":  result,
	})
}

func ResponseWithError(
	ctx *gin.Context,
	code int,
	message string,
	errors interface{},
) {
	ctx.JSON(code, gin.H{
		"message": message,
		"success": false,
		"errors":  errors,
	})
}
