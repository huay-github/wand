package models

import "time"

// Permission represents a API permission.
type Permission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

// Repository represents a API repository.
type Repository struct {
	ID            int64       `json:"id"`
	Owner         *User       `json:"owner"`
	Name          string      `json:"name"`
	FullName      string      `json:"full_name"`
	Description   string      `json:"description"`
	Private       bool        `json:"private"`
	Fork          bool        `json:"fork"`
	Parent        *Repository `json:"parent"`
	Empty         bool        `json:"empty"`
	Mirror        bool        `json:"mirror"`
	Size          int64       `json:"size"`
	HTMLURL       string      `json:"html_url"`
	SSHURL        string      `json:"ssh_url"`
	CloneURL      string      `json:"clone_url"`
	Website       string      `json:"website"`
	Stars         int         `json:"stars_count"`
	Forks         int         `json:"forks_count"`
	Watchers      int         `json:"watchers_count"`
	OpenIssues    int         `json:"open_issues_count"`
	DefaultBranch string      `json:"default_branch"`
	Created       time.Time   `json:"created_at"`
	Updated       time.Time   `json:"updated_at"`
	Permissions   *Permission `json:"permissions,omitempty"`
}
