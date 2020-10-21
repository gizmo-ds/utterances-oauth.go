package authorize

import (
	"net/http"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin/binding"
	"uapi/oauth"
	"uapi/state"
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
		ginx.Error(err)
	}

	http.Redirect(w, r, oauth.AuthURL(_state), http.StatusFound)
}
