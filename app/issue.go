package app

import (
	"errors"
	"log"
	"net/http"
	"strings"

	github_api "uapi/github-api"
	"uapi/oauth"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin"
)

func IssueHandler(c *gin.Context) {
	var form struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	ginx.ShouldBindJSON(c, &form)

	authorization := strings.Fields(c.GetHeader("Authorization"))
	if len(authorization) != 2 {
		ginx.R(http.StatusBadRequest, errors.New(`"Authorization" header is required`))
	}

	client := oauth.Client(c, authorization[1])
	if statusCode, err := github_api.User(client); err != nil || statusCode != http.StatusOK {
		ginx.R(http.StatusUnauthorized, errors.New("unauthorized"))
	}

	issue := github_api.Issue{
		Owner:  c.Param("owner"),
		Repo:   c.Param("repo"),
		Title:  form.Title,
		Body:   form.Body,
		Labels: []string{},
	}
	label := c.Query("label")
	if label != "" {
		issue.Labels = []string{label}
	}

	body, code, header, err := github_api.CreateIssue(client, issue)
	if err != nil {
		log.Println(err)
		ginx.R(http.StatusServiceUnavailable, errors.New("unable to post issue to GitHub"))
	}

	c.Header("Content-Type", header.Get("Content-Type"))
	c.Header("Content-Length", header.Get("Content-Length"))
	c.String(code, "%s", string(body))
}
