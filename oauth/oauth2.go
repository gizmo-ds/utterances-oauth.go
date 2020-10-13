package oauth

import (
	"context"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
	}
}

func AuthURL(state string) string {
	return getConfig().AuthCodeURL(state)
}

func AccessToken(ctx context.Context, code string) (token *oauth2.Token, err error) {
	return getConfig().Exchange(ctx, code)
}

func Client(ctx context.Context, accessToken string) *http.Client {
	return getConfig().Client(ctx, &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	})
}
