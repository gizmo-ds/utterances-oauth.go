package app

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"uapi/oauth"
	"uapi/state"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin"
)

func AuthorizedHandler(c *gin.Context) {
	var query struct {
		Code  string `form:"code" binding:"required"`
		State string `form:"state" binding:"required"`
	}
	ginx.ShouldBindQuery(c, &query)

	_state, err := state.DecryptState(query.State)
	if err != nil {
		ginx.R(http.StatusBadRequest, errors.New(`"state" is invalid`))
	} else if time.Now().Unix() > _state.Expires {
		ginx.R(http.StatusBadRequest, errors.New(`"state" is expired`))
	}

	token, err := oauth.AccessToken(c, query.Code)
	if err != nil {
		log.Println(err)
		ginx.R(http.StatusServiceUnavailable, errors.New("unable to load token from GitHub"))
	}

	_url, err := url.Parse(_state.Value)
	if err != nil {
		ginx.Error(err)
	}

	values := _url.Query()
	values.Set("utterances", token.AccessToken)
	_url.RawQuery = values.Encode()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    url.QueryEscape(token.AccessToken),
		Path:     "/token",
		Secure:   true,
		HttpOnly: true,
		Expires:  token.Expiry,
		SameSite: http.SameSiteNoneMode,
	})

	c.Redirect(http.StatusFound, _url.String())
}
