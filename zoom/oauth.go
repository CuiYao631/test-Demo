package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

const (
	clientID     = "Hv_kjoEnRcmU7hAo8jDyw"
	clientSecret = "z5NyvDjRmDqmHAZylAzyqlPK7NNBD2zQ"
	redirectURL  = "http://localhost:8080/callback"
)

var (
	authEndpoint  = "https://zoom.us/oauth/authorize"
	tokenEndpoint = "https://zoom.us/oauth/token"
)

// ZoomTokenResponse 包含Zoom API的访问令牌响应结构
type ZoomTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getAuthorizationURL() string {
	// 构建Zoom OAuth 2.0认证链接
	authURL := fmt.Sprintf("%s?client_id=%s&response_type=code&redirect_uri=%s", authEndpoint, clientID, redirectURL)
	return authURL
}
func main() {

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://zoom.us/oauth/authorize",
			TokenURL: "https://zoom.us/oauth/token",
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		authURL := getAuthorizationURL()
		fmt.Printf("请在浏览器中打开以下链接进行授权:\n%s\n", authURL)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			log.Fatal("无法获取访问令牌:", err)
		}

		// 这里可以使用 token 进行访问令牌的相关操作
		fmt.Println("访问令牌:", token.AccessToken)

	})

	fmt.Println("请在浏览器中访问 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
