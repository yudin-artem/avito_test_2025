package models

type PRStatus string

const (
	PRStatusOpen   PRStatus = "OPEN"
	PRStatusMerged PRStatus = "MERGED"
)

type Team struct {
	TeamName string `gorm:"primaryKey;type:varchar(255)" json:"team_name"`
	Members  []User `gorm:"foreignKey:TeamName;references:TeamName" json:"members"`
}

type User struct {
	UserID   string `gorm:"primaryKey;type:varchar(255)" json:"user_id"`
	Username string `gorm:"not null;type:varchar(255)" json:"username"`
	TeamName string `gorm:"not null;type:varchar(255);index" json:"team_name"`
	IsActive bool   `gorm:"default:true;index" json:"is_active"`

	Team Team `gorm:"foreignKey:TeamName;references:TeamName" json:"-"`
}

type PullRequest struct {
	PullRequestID      string   `gorm:"primaryKey;type:varchar(255)" json:"pull_request_id"`
	PullRequestName    string   `gorm:"not null;type:varchar(255)" json:"pull_request_name"`
	AuthorID           string   `gorm:"not null;type:varchar(255);index" json:"author_id"`
	Status             PRStatus `gorm:"type:pr_status;default:'OPEN';index" json:"status"`

	Reviewers PrReviewers `gorm:"foreignKey:PullRequestID;references:PullRequestID" json:"-"`
	Author    User                 `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
}

type PrReviewers struct {
	PullRequestID        string   `gorm:"primaryKey;type:varchar(255);column:pr_id" json:"pr_id"`
	Reviewer1ID string   `gorm:"type:varchar(255);column:reviewer_1_id" json:"reviewer_1_id"`
	Reviewer2ID string   `gorm:"type:varchar(255);column:reviewer_2_id" json:"reviewer_2_id"`

	Reviewer1 User `gorm:"foreignKey:Reviewer1ID;references:UserID" json:"-"`
	Reviewer2 User `gorm:"foreignKey:Reviewer2ID;references:UserID" json:"-"`
}