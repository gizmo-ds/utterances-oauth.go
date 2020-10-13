package github_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type (
	Issue struct {
		Owner  string   `json:"-"`
		Repo   string   `json:"-"`
		Labels []string `json:"labels,omitempty'"`
		Title  string   `json:"title"`
		Body   string   `json:"body"`
	}
)

func User(client *http.Client) (statusCode int, err error) {
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	return
}

func CreateIssue(client *http.Client, issue Issue) (body []byte, code int, header http.Header, err error) {
	_url, _ := url.Parse("https://api.github.com/")
	_url.Path = fmt.Sprintf("/repos/%s/%s/issues", issue.Owner, issue.Repo)

	reqBody, err := json.Marshal(&issue)
	if err != nil {
		return
	}

	resp, err := client.Post(_url.String(), "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	code = resp.StatusCode
	header = resp.Header
	body, err = ioutil.ReadAll(resp.Body)
	return
}
