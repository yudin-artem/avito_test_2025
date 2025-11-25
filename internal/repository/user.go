package repository

import (
	"github.com/yudin-artem/avito_test_2025/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, models.ErrNotFound
	}

	return &user, nil
}

func (r *UserRepository) SetIsActive(req *models.SetActiveRequest) (*models.User, error) {
	user, err := r.GetUser(req.UserID)
	if err != nil {
		return nil, err
	}

	user.IsActive = req.IsActive
	if err := r.db.Model(user).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetReview(userID string) (*models.UserReviewsResponse, error) {
	_, err := r.GetUser(userID)
	if err != nil {
		return nil, err
	}

	var reviewers []models.PrReviewers
	if err := r.db.Where("reviewer_1_id = ? OR reviewer_2_id = ?", userID, userID).Find(&reviewers).Error; err != nil {
		return nil, err
	}

	var prID []string
	for _, assignment := range reviewers {
		prID = append(prID, assignment.PullRequestID)
	}

	var prs []models.PullRequest
	if err := r.db.Where("pull_request_id IN ?", prID).Find(&prs).Error; err != nil {
		return nil, err
	}

	prsh := make([]models.PullRequestShort, 0)
	for _, pr := range prs {
		prsh = append(prsh, models.PullRequestShort{
			PullRequestID:   pr.PullRequestID,
			PullRequestName: pr.PullRequestName,
			AuthorID:        pr.AuthorID,
			Status:          string(pr.Status),
		})
	}

	return &models.UserReviewsResponse{
		UserID: userID,
		PullRequests:    prsh,
	}, nil
}