package service

import (
	"github.com/yudin-artem/avito_test_2025/internal/repository"
	"github.com/yudin-artem/avito_test_2025/internal/models"
)

type PRService struct {
	repo       repository.PR
}

func NewPRService(repo repository.PR) *PRService {
	return &PRService{
		repo:repo,
	}
}

func (s *PRService) CreatePR(req *models.CreatePR) (*models.PullRequestResponse, error) {
	reviewers, err := s.repo.FindReviewersByAuthorID(req.AuthorID)
	if err != nil {
		return nil, err
	}

	var prReviewers = models.PrReviewers{
		PullRequestID: req.PullRequestID,
	}

	prReviewers.Reviewer1ID = (*reviewers)[0]
	if len(*reviewers) == 2 {
		prReviewers.Reviewer2ID = (*reviewers)[1]
	}

	var pr = models.PullRequest{
		PullRequestID: req.PullRequestID,
		PullRequestName: req.PullRequestName,
		AuthorID: req.AuthorID,

		Reviewers: prReviewers,
	}

	pullRequest, err := s.repo.CreatePR(&pr)
	if err != nil {
		return nil, err
	}

	resp := models.PullRequestResponse{
		PullRequestID: pullRequest.PullRequestID,
		PullRequestName: pullRequest.PullRequestName,
		AuthorID: pullRequest.AuthorID,
		Status: string(pullRequest.Status),
		Reviewers: []string{pullRequest.Reviewers.Reviewer1ID, pullRequest.Reviewers.Reviewer2ID},
	}

	return &resp, nil
}

func (s *PRService) MergePR(req *models.MergePR) (*models.PullRequestResponse, error) {
	pullRequest, err := s.repo.MergePR(req.PullRequestID)
	if err != nil {
		return nil, err
	}

	resp := models.PullRequestResponse{
		PullRequestID: pullRequest.PullRequestID,
		PullRequestName: pullRequest.PullRequestName,
		AuthorID: pullRequest.AuthorID,
		Status: string(pullRequest.Status),
		Reviewers: []string{pullRequest.Reviewers.Reviewer1ID, pullRequest.Reviewers.Reviewer2ID},
	}

	return &resp, nil
}

func (s *PRService) Reassign(req *models.ReassignRequest) (*models.ReassignResponse, error)  {
	pullRequest, id, err := s.repo.Reassign(req.PullRequestID, req.OldUserID)
	if err != nil {
		return nil, err
	}

	var res = models.ReassignResponse{
		PR: &models.PullRequestResponse{
				PullRequestID: pullRequest.PullRequestID,
				PullRequestName: pullRequest.PullRequestName,
				AuthorID: pullRequest.AuthorID,
				Status: string(pullRequest.Status),
				Reviewers: []string{pullRequest.Reviewers.Reviewer1ID, pullRequest.Reviewers.Reviewer2ID},
			},
		ReplacedBy: id,
	}

	return &res, nil
}

