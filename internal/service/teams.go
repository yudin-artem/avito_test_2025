package service

import (
	"github.com/yudin-artem/avito_test_2025/internal/repository"
	"github.com/yudin-artem/avito_test_2025/internal/models"
)

type TeamService struct {
	repo       repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{
		repo:repo,
	}
}

func (s *TeamService) AddTeam(req *models.AddTeamRequest) (*models.Team, error) {
	team, err := s.repo.AddTeam(req)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) GetTeam(team_name string) (*models.Team, error) {
	team, err := s.repo.GetTeam(team_name)
	if err != nil {
		return nil, err
	}

	return team, nil
}

