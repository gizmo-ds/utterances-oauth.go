package issue

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/mux"
	"uapi/cors"
	github_api "uapi/github-api"
	"uapi/oauth"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.Add(w, r)
	var owner, repo string
	vars := mux.Vars(r)
	if vars != nil {
		owner = vars["owner"]
		repo = vars["repo"]
	} else {
		reg, err := regexp.Compile(`\/repos\/([^/]*)\/([^/]*)\/issues`)
		if err != nil {
			http.Error(w, "BadRequest", http.StatusBadRequest)
			return
		}
		arr := reg.FindAllStringSubmatch(r.URL.Path, -1)
		owner = arr[0][1]
		repo = arr[0][2]
	}

	var info struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	if err := binding.JSON.Bind(r, &info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorization := strings.Fields(r.Header.Get("Authorization"))
	if len(authorization) != 2 {
		ginx.R(http.StatusBadRequest, errors.New(`"Authorization" header is required`))
	}

	ctx := context.Background()
	client := oauth.Client(ctx, authorization[1])
	if statusCode, err := github_api.User(client); err != nil || statusCode != http.StatusOK {
		ginx.R(http.StatusUnauthorized, errors.New("unauthorized"))
	}

	issue := github_api.Issue{
		Owner:  owner,
		Repo:   repo,
		Title:  info.Title,
		Body:   info.Body,
		Labels: []string{},
	}
	label := r.URL.Query().Get("label")
	if label != "" {
		issue.Labels = []string{label}
	}

	body, code, header, err := github_api.CreateIssue(client, issue)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to post issue to GitHub", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", header.Get("Content-Type"))
	w.Header().Add("Content-Length", header.Get("Content-Length"))
	_, _ = fmt.Fprint(w, body)
}
