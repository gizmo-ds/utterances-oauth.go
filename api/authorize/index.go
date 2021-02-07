package authorize

import (
	"net/http"

	"uapi/oauth"
	"uapi/state"

	"github.com/gin-gonic/gin/binding"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RedirectUri string `form:"redirect_uri" binding:"required"`
	}
	if err := binding.Query.Bind(r, &query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_state, err := state.EncryptState(query.RedirectUri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, oauth.AuthURL(_state), http.StatusFound)
}
