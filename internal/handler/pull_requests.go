package handler

import (
	"net/http"

	"github.com/yudin-artem/avito_test_2025/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPR(c *gin.Context) {
	var req models.CreatePR
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", err.Error()))
		return
	}

	pullRequest, err := h.services.PR.CreatePR(&req)
	if err != nil {
		if err == models.ErrPRExists {
			c.JSON(http.StatusConflict, createErrorResponse("PR_EXISTS", err.Error()))
		} else if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pr": pullRequest})
}

func (h *Handler) mergePR(c *gin.Context) {
	var req models.MergePR
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", err.Error()))
		return
	}

	pullRequest, err := h.services.PR.MergePR(&req)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"pr": pullRequest})
}

func (h *Handler) reassign(c *gin.Context) {
	var req models.ReassignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", err.Error()))
		return
	}

	res, err := h.services.PR.Reassign(&req)
	if  err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else if err == models.ErrPRMerged {
			c.JSON(http.StatusConflict, createErrorResponse("PR_MERGED", err.Error()))
		} else if err == models.ErrNotAssigned {
			c.JSON(http.StatusConflict, createErrorResponse("NOT_ASSIGNED", err.Error()))
		} else if err == models.ErrNoCandidate {
			c.JSON(http.StatusConflict, createErrorResponse("NO_CANDIDATE", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, res)
}