package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					log.Println(e)
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
				return
			}
		}()
		c.Next()
	}
}
