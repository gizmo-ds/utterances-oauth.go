package oauth

import (
	"context"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var config *oauth2.Config

func init() {
	config = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
	}
}

func AuthURL(state string) string {
	return config.AuthCodeURL(state)
}

func AccessToken(ctx context.Context, code string) (token *oauth2.Token, err error) {
	return config.Exchange(ctx, code)
}

func Client(ctx context.Context, accessToken string) *http.Client {
	return config.Client(ctx, &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	})
}
