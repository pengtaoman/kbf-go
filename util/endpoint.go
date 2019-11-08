package util

import (
	"github.com/gin-gonic/gin"

)

func ReadEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong1111111",
	})
	//strings.Join()
}
