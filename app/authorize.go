package app

import (
	"net/http"

	"uapi/oauth"
	"uapi/state"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler(c *gin.Context) {
	var query struct {
		RedirectUri string `form:"redirect_uri" binding:"required"`
	}
	ginx.ShouldBindQuery(c, &query)

	_state, err := state.EncryptState(query.RedirectUri)
	if err != nil {
		ginx.Error(err)
	}

	c.Redirect(http.StatusFound, oauth.AuthURL(_state))
}
