package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL is URL of github api for issue
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is result of issue search
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue of search result
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Labels    []*Label
	Assignees []*User
	Milestone *Milestone
}

// User of Issue
type User struct {
	Login     string
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
}

// Label of Issue
type Label struct {
	Name  string
	Color string
}

// Milestone of Issue
type Milestone struct {
	HTMLURL      string `json:"html_url"`
	Number       int
	State        string
	Title        string
	Description  string
	Creator      *User
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}

// SearchIssues queries the GitHub issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
