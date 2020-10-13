package app

import (
	"net/http"
	"uapi/oauth"
	"uapi/state"

	"github.com/gin-gonic/gin"
)

func AuthorizeHandler(c *gin.Context) {
	var query struct {
		RedirectUri string `form:"redirect_uri" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_state, err := state.EncryptState(query.RedirectUri)
	if err != nil {
		panic(err)
	}

	c.Redirect(http.StatusFound, oauth.AuthURL(_state))
}
