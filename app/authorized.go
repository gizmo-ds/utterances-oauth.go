package app

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"uapi/oauth"
	"uapi/state"

	"github.com/gin-gonic/gin"
)

func AuthorizedHandler(c *gin.Context) {
	var query struct {
		Code  string `form:"code" binding:"required"`
		State string `form:"state" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_state, err := state.DecryptState(query.State)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `"state" is invalid`,
		})
		return
	} else if time.Now().Unix() > _state.Expires {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": `"state" is expired`,
		})
		return
	}

	token, err := oauth.AccessToken(c, query.Code)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to load token from GitHub",
		})
		return
	}

	_url, err := url.Parse(_state.Value)
	if err != nil {
		panic(err)
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
