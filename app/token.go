package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenHandler(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, token)
}
