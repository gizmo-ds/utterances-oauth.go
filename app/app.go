package app

import (
	"net/http"
	"os"
	"strings"
	"time"

	"uapi/app/middleware"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Start() {
	app := RegisterRouter()
	log.Fatal(app.Run(":5000"))
}

func RegisterRouter() *gin.Engine {
	app := gin.New()
	app.Use(middleware.Error())
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(os.Getenv("ORIGINS"), ","),
		AllowMethods: []string{http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{
			"X-Requested-With", "X-HTTP-Method-Override",
			"Content-Type", "Accept",
			"Authorization", "label",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "alive")
	})
	app.POST("/token", TokenHandler)
	app.GET("/authorize", AuthorizeHandler)
	app.GET("/authorized", AuthorizedHandler)
	app.POST("/repos/:owner/:repo/issues", IssueHandler)
	return app
}
