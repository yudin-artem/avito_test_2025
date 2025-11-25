CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE teams (
    team_name VARCHAR(255) PRIMARY KEY
);

CREATE TABLE users (
    user_id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    team_name VARCHAR(255) NOT NULL REFERENCES teams(team_name),
    is_active BOOLEAN DEFAULT true
);

CREATE TABLE pull_requests (
    pull_request_id VARCHAR(255) PRIMARY KEY,
    pull_request_name VARCHAR(255) NOT NULL,
    author_id VARCHAR(255) NOT NULL REFERENCES users(user_id),
    status pr_status DEFAULT 'OPEN'
);

CREATE TABLE pr_reviewers (
    pr_id VARCHAR(255) PRIMARY KEY REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    reviewer_1_id VARCHAR(255) REFERENCES users(user_id),
    reviewer_2_id VARCHAR(255) REFERENCES users(user_id)
);

CREATE INDEX idx_users_team_name ON users(team_name);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_pull_requests_status ON pull_requests(status);
CREATE INDEX idx_pull_requests_author ON pull_requests(author_id);
CREATE INDEX idx_pr_reviewers_reviewer_1 ON pr_reviewers(reviewer_1_id);
CREATE INDEX idx_pr_reviewers_reviewer_2 ON pr_reviewers(reviewer_2_id);