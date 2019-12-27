package github

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// 许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息。为了说明这一点，我们通过Github的issue查询服务来演示类似的用法。

const IssuesURL = "https://api.github.com/search/issue"

type IssuesSearchResult struct {
	TotalCount		int		`json:"total_count"`
	Items			[]*Issue
}

type Issue struct {
	Number			int
	HTMLURL			string		`json:"html_url"`
	Title			string
	State			string
	User			*User
	CreateAt		time.Time	`json:created_at`
	Body			string
}

type User struct {
	Login		string
	HTMLURL		string		`json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, ","))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}