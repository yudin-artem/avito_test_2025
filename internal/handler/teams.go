package handler

import (
	"net/http"

	"github.com/yudin-artem/avito_test_2025/internal/models"
	
	"github.com/gin-gonic/gin"
)

func (h* Handler) addTeam(c *gin.Context) {
	var team *models.AddTeamRequest
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", err.Error()))
		return
	}

	createdTeam, err := h.services.Team.AddTeam(team)
	if err != nil {
		if err == models.ErrTeamExists {
			c.JSON(http.StatusBadRequest, createErrorResponse("TEAM_EXISTS", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"team": createdTeam})
}

func (h* Handler) getTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		c.JSON(http.StatusBadRequest, createErrorResponse("INVALID_INPUT", "team_name is required"))
		return
	}

	team, err := h.services.Team.GetTeam(teamName)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, createErrorResponse("NOT_FOUND", err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, createErrorResponse("INTERNAL_ERROR", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, team)
}

