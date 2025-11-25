package repository

import (
	"errors"

	"github.com/yudin-artem/avito_test_2025/internal/models"
	
	"gorm.io/gorm"
)

type PRRepository struct {
	db *gorm.DB
}

func NewPRRepository(db *gorm.DB) *PRRepository {
	return &PRRepository{db: db}
}

func (r *PRRepository) getPR(PullRequestID string) (*models.PullRequest, error) {
	var pullRequest models.PullRequest
	if err := r.db.Preload("Reviewers").Where("pull_request_id = ?", PullRequestID).First(&pullRequest).Error; err != nil {
		return nil, models.ErrNotFound
	}

	return &pullRequest, nil
}

func (r *PRRepository) CreatePR(pr *models.PullRequest) (*models.PullRequest, error) {
	var existingPR models.PullRequest
	if err := r.db.Where("pull_request_id = ?", pr.PullRequestID).First(&existingPR).Error; err == nil {
		return nil, models.ErrPRExists
	}

	if err := r.db.Create(pr).Error; err != nil {
		return nil, err
	}

	return pr, nil
}

func (r *PRRepository) FindReviewersByAuthorID(authorID string) (*[]string , error) {
	var author models.User
	if err := r.db.Where("user_id = ?", authorID).First(&author).Error; err != nil {
		return nil, models.ErrNotFound
	}

	var teamMembers []models.User
	if err := r.db.Where("team_name = ? AND user_id != ? AND is_active = ?", author.TeamName, authorID, true).Find(&teamMembers).Error; err != nil {
		return nil, err
	}

	if len(teamMembers) == 0 {
		return nil, models.ErrNoCandidate
	}

	reviewerCount := min(2, len(teamMembers))
	reviewers := make([]string, reviewerCount)
	for i := 0; i < reviewerCount; i++ {
		reviewers[i] = teamMembers[i].UserID
	}

	return &reviewers, nil
}

func (r *PRRepository) MergePR(prID string) (*models.PullRequest ,error) {
	pullRequest, err := r.getPR(prID)
	if err != nil {
		return nil, err
	}

	pullRequest.Status = models.PRStatusMerged

	if err := r.db.Model(pullRequest).Where("pull_request_id = ?", pullRequest.PullRequestID).Updates(pullRequest).Error; err != nil {
		return nil, err
	}

	return pullRequest, nil
}

func (r *PRRepository) Reassign(prID string, oldReviewerID string) (*models.PullRequest, string, error)  {
	pullRequest, err := r.getPR(prID)
	if err != nil {
		return nil, "", err
	}

	if pullRequest.Status == "MERGED" {
		return nil, "", models.ErrPRMerged
	}

	var reviewers models.PrReviewers
	if err := r.db.Where("pr_id = ?", prID).First(&reviewers).Error; err != nil {
		return nil, "", err
	}

	if reviewers.Reviewer1ID != oldReviewerID && reviewers.Reviewer2ID != oldReviewerID {
		return nil, "", models.ErrNotAssigned
	}

	var author models.User
	if err := r.db.Where("user_id = ?", pullRequest.AuthorID).First(&author).Error; err != nil {
		return nil, "", err
	}
	
	var replacement models.User
	query := r.db.Where("team_name = ? AND user_id != ? AND is_active = ?", 
		author.TeamName, pullRequest.AuthorID, true)
	
	if reviewers.Reviewer1ID != "" {
		query = query.Where("user_id != ?", reviewers.Reviewer1ID)
	}
	if reviewers.Reviewer2ID != "" {
		query = query.Where("user_id != ?", reviewers.Reviewer2ID)
	}

	if err := query.First(&replacement).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", models.ErrNoCandidate
		}
		return nil, "", err
	}


	if reviewers.Reviewer1ID == oldReviewerID {
		reviewers.Reviewer1ID = replacement.UserID
	} else {
		reviewers.Reviewer2ID = replacement.UserID
	}

	pullRequest.Reviewers = reviewers

	if err := r.db.Model(pullRequest).Updates(pullRequest).Error; err != nil {
			return nil, "", err
	}

	return pullRequest, replacement.UserID, nil
}
