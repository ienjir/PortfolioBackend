package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/PortfolioBackend/internal/controller"
)

func Routes(router *gin.Engine) {
	userRoutes := router.Group("/calctime")
	{
		userRoutes.GET("/", controller.CalcTime)
	}
}
