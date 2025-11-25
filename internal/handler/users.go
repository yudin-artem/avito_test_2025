package handler

import (
	"net/http"

	"github.com/yudin-artem/avito_test_2025/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) setIsActive(c *gin.Context) {
	var req *models.SetActiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", err.Error()))
		return
	}

	user, err := h.services.User.SetIsActive(req)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}


func (h *Handler) getReview(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", "user_id is required"))
		return
	}

	res, err := h.services.User.GetReview(userID)
	if err != nil {
		if(err == models.ErrNotFound) {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, res)
}

