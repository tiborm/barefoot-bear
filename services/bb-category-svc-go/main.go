package main

import (
	"fmt"
	"net/http"

	"math/rand"

	"golang.org/x/oauth2"
)

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

func main() {
	oauthStateString = RandStringBytes(64)
	fmt.Println(oauthStateString)

    http.HandleFunc("/", handleMain)
    http.HandleFunc("/login", handleGoogleLogin)
    http.HandleFunc("/callback", handleGoogleCallback)
    fmt.Println("Server started at http://localhost:5491")
    http.ListenAndServe(":5491", nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
    var htmlIndex = `<html><body><a href="/login">Google Log In</a></body></html>`
    w.Write([]byte(htmlIndex))
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    URL := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
    http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("you did great"))
}