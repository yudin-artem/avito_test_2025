package repository

import(
	"github.com/yudin-artem/avito_test_2025/internal/models"
	"gorm.io/gorm"
)

type User interface {
	GetUser(string) (*models.User, error)
	SetIsActive(*models.SetActiveRequest) (*models.User, error)
	GetReview(string) (*models.UserReviewsResponse, error)
}

type Team interface {
	getActiveMembers(string) (*[]models.User, error)
	AddTeam(*models.AddTeamRequest) (*models.Team, error)
	GetTeam(string) (*models.Team, error)
}

type PR interface {
	getPR(string) (*models.PullRequest, error)
	CreatePR(*models.PullRequest) (*models.PullRequest ,error)
	MergePR(string) (*models.PullRequest ,error)
	Reassign(string, string) (*models.PullRequest, string, error)
	FindReviewersByAuthorID(string) (*[]string, error)
}

type Repository struct {
	User
	Team
	PR
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Team: NewTeamRepository(db),
		PR:   NewPRRepository(db),
	}
}