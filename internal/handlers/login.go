package handlers

import (
	"net/http"

	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"
)

// TODO refactor this to use a config file or environment variables
var (
	oauthConfig = &oauth2.Config{
        RedirectURL:  "http://localhost:5491/callback",
        ClientID:     "YOUR_AMAZON_CLIENT_ID",
        ClientSecret: "YOUR_AMAZON_CLIENT_SECRET",
        Scopes:       []string{"profile", "postal_code"},
        Endpoint: oauth2.Endpoint{
            AuthURL:  "https://www.amazon.com/ap/oa",
            TokenURL: "https://api.amazon.com/auth/o2/token",
        },
    }
	oauthStateString = ""
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func Login() {
	oauthStateString = RandStringBytes(16)
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    URL := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
    http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("you did great"))
}
