package search

import (
	"encoding/json"
	"fmt"
	"go_code/go_exercise/go_programming_4/issues/github"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*github.IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	if q == "" {
		q = "repo:golang/go+is:open+json+decoder"
	}
	resp, err := http.Get(github.IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result github.IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
