package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataController struct {
}

func NewDataController() *DataController {
	return &DataController{}
}

func (c *DataController) GetData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get data by ID",
	})
}

func (c *DataController) AddData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data added successfully",
	})
}

func (c *DataController) UpdateData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data updated successfully",
	})
}
