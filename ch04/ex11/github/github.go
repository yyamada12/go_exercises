package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// IssueURL is URL of github api for issue
const IssueURL = "https://api.github.com/repos/%s/%s/issues"

// IssueOperator gets, creates, updates and closes github issue
type IssueOperator struct {
	token string
	url   string
}

// IssueResult is result from api
type IssueResult struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User is attribute of IssueResult
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// NewIssueOperator returns new IssueOperator from github token, owner and repo
func NewIssueOperator(token, owner, repo string) IssueOperator {
	url := fmt.Sprintf(IssueURL, url.QueryEscape(owner), url.QueryEscape(repo))
	return IssueOperator{token, url}
}

// GetIssue gets issue by issue_number
func (i IssueOperator) GetIssue(number string) (*IssueResult, error) {
	getIssueURL := i.url + "/" + url.QueryEscape(number)
	resp, err := http.Get(getIssueURL)
	if err != nil {
		return nil, err
	}
	return handleResp(resp, http.StatusOK, "get issue failed: %s")
}

type issuePost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreateIssue posts a issue to the GitHub
func (i IssueOperator) CreateIssue(title, body string) (*IssueResult, error) {
	param, err := json.Marshal(issuePost{title, body})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", i.url, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+i.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return handleResp(resp, http.StatusCreated, "post issue failed: %s")
}

// UpdateIssue updates issue by issue_number
func (i IssueOperator) UpdateIssue(number, title, body string) (*IssueResult, error) {
	updateIssueURL := i.url + "/" + url.QueryEscape(number)

	param, err := json.Marshal(issuePost{title, body})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", updateIssueURL, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+i.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return handleResp(resp, http.StatusOK, "update issue failed: %s")
}

// CloseIssue closes issue by issue_number
func (i IssueOperator) CloseIssue(number string) (*IssueResult, error) {
	closeIssueURL := i.url + "/" + url.QueryEscape(number)

	param, err := json.Marshal(struct {
		State string `json:"state"`
	}{"closed"})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", closeIssueURL, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+i.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return handleResp(resp, http.StatusOK, "close issue failed: %s")
}

func handleResp(resp *http.Response, expectedStatus int, errMsg string) (*IssueResult, error) {
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return nil, fmt.Errorf(errMsg, resp.Status)
	}

	var result IssueResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &result, nil
}
