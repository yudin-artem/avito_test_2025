package handler

import (
	"fmt"

	"github.com/yudin-artem/avito_test_2025/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s %s %d %s\n",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	users := router.Group("/users")
	{
		users.POST("/setIsActive", h.setIsActive)
		users.GET("/getReview", h.getReview)
	}
	teams := router.Group("/team")
	{
		teams.POST("/add", h.addTeam)
		teams.GET("/get", h.getTeam)
	}
	pullRequest := router.Group("/pullRequest")
	{
		pullRequest.POST("/create", h.createPR)
		pullRequest.POST("/merge", h.mergePR)
		pullRequest.POST("/reassign", h.reassign)
	}

	return router
}

