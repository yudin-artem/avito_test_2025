package service

import (
	"github.com/yudin-artem/avito_test_2025/internal/models"
	"github.com/yudin-artem/avito_test_2025/internal/repository"
)


type User interface{
	SetIsActive(*models.SetActiveRequest) (*models.User, error)
	GetReview(string) (*models.UserReviewsResponse, error)
}

type Team interface {
	AddTeam(*models.AddTeamRequest) (*models.Team, error)
	GetTeam(string) (*models.Team, error)
}

type PR interface {
	CreatePR(*models.CreatePR) (*models.PullRequestResponse, error)
	MergePR(*models.MergePR) (*models.PullRequestResponse, error)
	Reassign(*models.ReassignRequest) (*models.ReassignResponse, error)
}

type Service struct {
	User
	Team
	PR
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Team: NewTeamService(repos.Team),
		PR:   NewPRService(repos.PR),
	}
}