package cors

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

var c = cors.New(cors.Options{
	AllowedHeaders: []string{"Origin"},
	AllowedOrigins: strings.Split(os.Getenv("ORIGINS"), ","),
	AllowedMethods: []string{http.MethodPost, http.MethodGet, http.MethodOptions},
	ExposedHeaders: []string{
		"X-Requested-With", "X-HTTP-Method-Override",
		"Content-Type", "Accept",
		"Authorization", "label",
	},
	AllowCredentials: true,
	MaxAge:           24 * 60,
})

func Add(w http.ResponseWriter, r *http.Request) {
	c.HandlerFunc(w, r)
}
