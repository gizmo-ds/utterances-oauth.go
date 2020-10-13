package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	github_api "uapi/github-api"
	"uapi/oauth"
)

func IssueHandler(c *gin.Context) {
	var form struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorization := strings.Fields(c.GetHeader("Authorization"))
	if len(authorization) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `"Authorization" header is required`,
		})
		return
	}

	client := oauth.Client(c, authorization[1])
	if statusCode, err := github_api.User(client); err != nil || statusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
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
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "unable to post issue to GitHub",
		})
		return
	}

	c.Header("Content-Type", header.Get("Content-Type"))
	c.Header("Content-Length", header.Get("Content-Length"))
	c.String(code, "%s", string(body))
}
