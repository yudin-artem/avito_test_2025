package models

import (
	"errors"
)

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type AddTeamRequest struct {
	TeamName string `json:"team_name"`
	Members []struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		IsActive bool   `json:"is_active"`
	} `json:"members"`
}

type SetActiveRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

type CreatePR struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

type MergePR struct {
	PullRequestID string `json:"pull_request_id"`
}

type ReassignRequest struct {
	PullRequestID string `json:"pull_request_id"`
	OldUserID     string `json:"old_user_id"`
}

type UserReviewsResponse struct {
	UserID        string             `json:"user_id"`
	PullRequests  []PullRequestShort `json:"pull_requests"`
}

type PullRequestShort struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
}

type PullRequestResponse struct {
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
	Status          string `json:"status"`
	Reviewers       []string `json:"assigned_reviewers"`
}

type ReassignResponse struct {
	PR          *PullRequestResponse `json:"pr"`
	ReplacedBy  string       `json:"replaced_by"`
}

var (
	ErrTeamExists    = errors.New("team name already exists")
	ErrPRExists      = errors.New("PR id already exists")
	ErrPRMerged      = errors.New("cannot reassign on merged PR")
	ErrNotAssigned   = errors.New("reviewer is not assigned to this PR")
	ErrNoCandidate   = errors.New("no active replacement candidate in team")
	ErrNotFound      = errors.New("resource not found")
)