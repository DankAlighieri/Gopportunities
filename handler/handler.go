package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func GetOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
