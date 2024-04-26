package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// Get Opening
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /v1/opening",
			})
		})

		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /v1/opening",
			})
		})

		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /v1/opening",
			})
		})

		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /v1/opening",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /v1/opening",
			})
		})
	}
}
