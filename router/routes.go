package router

import (
	"github.com/MarcosSmeets/GoOportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// Get Opening
		v1.GET("/opening", handler.ShowOpningHandler)

		v1.POST("/opening", handler.CreateOpningHandler)

		v1.DELETE("/opening", handler.DeleteOpningHandler)

		v1.PUT("/opening", handler.UpdateOpningHandler)

		v1.GET("/openings", handler.ListOpningHandler)
	}
}
