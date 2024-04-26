package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET /v1/opening",
	})
}
