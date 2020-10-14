package app

import (
	"net/http"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin"
)

func TokenHandler(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		ginx.R(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, token)
}
