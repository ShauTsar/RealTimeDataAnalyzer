package routes

import (
	"RealTimeDataAnalyzer/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	dataController := controllers.NewDataController()

	router.GET("/data/:id", dataController.GetData)
	router.POST("/data", dataController.AddData)
	router.PUT("/data/:id", dataController.UpdateData)
}
