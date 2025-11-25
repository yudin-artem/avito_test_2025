package service

import (
	"github.com/yudin-artem/avito_test_2025/internal/repository"
	"github.com/yudin-artem/avito_test_2025/internal/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo:repo,
	}
}

func (s *UserService) SetIsActive(req *models.SetActiveRequest) (*models.User, error) {
	user, err := s.repo.SetIsActive(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetReview(team_name string) (*models.UserReviewsResponse, error) {
	team, err := s.repo.GetReview(team_name)
	if err != nil {
		return nil, err
	}

	return team, nil
}

