package handler

import (
    "net/http"
	
    "github.com/yudin-artem/avito_test_2025/internal/models"

    "github.com/gin-gonic/gin"
)

func createErrorResponse(code, message string) models.ErrorResponse {
    var resp models.ErrorResponse
    resp.Error.Code = code
    resp.Error.Message = message
    return resp
}

func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
    })
}