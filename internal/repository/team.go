package repository

import (
	"github.com/yudin-artem/avito_test_2025/internal/models"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) getActiveMembers(teamName string) (*[]models.User, error) {
	var members []models.User
	if err := r.db.Where("team_name = ? AND is_active = ?", teamName, true).Find(&members).Error; err != nil {
		return nil, err
	}

	return &members, nil
}

func (r *TeamRepository) AddTeam(team *models.AddTeamRequest) (*models.Team, error) {
	var existingTeam models.Team
	if err := r.db.Where("team_name = ?", team.TeamName).First(&existingTeam).Error; err == nil {
		return nil, models.ErrTeamExists
	}

	tx := r.db.Begin()
	defer func() {
		if s := recover(); s != nil {
			tx.Rollback()
		}
	}()

	newTeam := models.Team{
		TeamName: team.TeamName,
		Members:  []models.User{},
	}

	if err := tx.Model(&models.Team{}).Create(newTeam).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for el := range team.Members {
		newUser := models.User{
			UserID:   team.Members[el].UserID,
			Username: team.Members[el].Username,
			TeamName: team.TeamName,
			IsActive: team.Members[el].IsActive,
		}

		if err := tx.Save(&newUser).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		newTeam.Members = append(newTeam.Members, newUser)
	}

	return  &newTeam, tx.Commit().Error
}

func (r *TeamRepository) GetTeam(team_name string) (*models.Team, error) {
	var team models.Team
	if err := r.db.Preload("Members").Where("team_name = ?", team_name).First(&team).Error; err != nil {
		return nil, models.ErrNotFound
	}

	return &team, nil
}