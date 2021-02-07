package authorized

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"uapi/oauth"
	"uapi/state"

	"github.com/gin-gonic/gin/binding"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var query struct {
		Code  string `form:"code" binding:"required"`
		State string `form:"state" binding:"required"`
	}
	if err := binding.Query.Bind(r, &query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_state, err := state.DecryptState(query.State)
	if err != nil {
		http.Error(w, `"state" is invalid`, http.StatusBadRequest)
		return
	} else if time.Now().Unix() > _state.Expires {
		http.Error(w, `"state" is expired`, http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	token, err := oauth.AccessToken(ctx, query.Code)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to load token from GitHub", http.StatusServiceUnavailable)
		return
	}

	_url, err := url.Parse(_state.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	values := _url.Query()
	values.Set("utterances", token.AccessToken)
	_url.RawQuery = values.Encode()

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    url.QueryEscape(token.AccessToken),
		Path:     "/token",
		Secure:   true,
		HttpOnly: true,
		Expires:  token.Expiry,
		SameSite: http.SameSiteNoneMode,
	})

	http.Redirect(w, r, _url.String(), http.StatusFound)
}
