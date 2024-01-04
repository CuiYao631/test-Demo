package main

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
//	"golang.org/x/oauth2"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//)
//
//const (
//	clientID     = "**************"
//	clientSecret = "**************"
//	redirectURI  = "https://localhost:8080/callback"
//)
//
//var (
//	authEndpoint  = "https://zoom.us/oauth/authorize"
//	tokenEndpoint = "https://zoom.us/oauth/token"
//)
//
//// ZoomTokenResponse 包含Zoom API的访问令牌响应结构
//type ZoomTokenResponse struct {
//	AccessToken string `json:"access_token"`
//	TokenType   string `json:"token_type"`
//	ExpiresIn   int    `json:"expires_in"`
//}
//
//func main() {
//	// 获取授权码
//	//authCode := getAuthorizationCode()
//
//	config := &oauth2.Config{
//		ClientID:     clientID,
//		ClientSecret: clientSecret,
//		RedirectURL:  redirectURI,
//		//Scopes:       []string{"user:email"},
//		Endpoint: oauth2.Endpoint{
//			AuthURL:  authEndpoint,
//			TokenURL: tokenEndpoint,
//		},
//	}
//
//	e := echo.New()
//
//	// 中间件，用于处理 OAuth 2.0 认证的回调
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	e.GET("/", func(c echo.Context) error {
//		url := config.AuthCodeURL("state")
//		return c.Redirect(http.StatusTemporaryRedirect, url)
//	})
//
//	e.GET("/callback", func(c echo.Context) error {
//		code := c.QueryParam("code")
//
//		token, err := config.Exchange(context.Background(), code)
//		if err != nil {
//			log.Println("无法获取访问令牌:", err)
//		}
//
//		// 这里可以使用 token 进行访问令牌的相关操作
//		fmt.Println("访问令牌:", token.AccessToken)
//
//		return c.String(http.StatusOK, "OAuth 2.0 认证成功")
//	})
//
//	fmt.Println("请在浏览器中访问 http://localhost:8080")
//	e.Start(":8080")
//
//	//// 使用授权码获取访问令牌
//	//token, err := getAccessToken(authCode)
//	//if err != nil {
//	//	fmt.Println("获取访问令牌时发生错误:", err)
//	//	return
//	//}
//	//
//	//fmt.Println("Zoom API访问令牌:", token.AccessToken)
//}
//
//func getAuthorizationCode() string {
//	// 构建授权码请求URL
//	authURL := fmt.Sprintf("%s?client_id=%s&response_type=code&redirect_uri=%s", authEndpoint, clientID, redirectURI)
//
//	fmt.Printf("请在浏览器中打开以下链接进行授权:\n%s\n", authURL)
//
//	// 从用户输入中获取授权码
//	fmt.Print("请输入授权码: ")
//	var authCode string
//	fmt.Scan(&authCode)
//
//	return authCode
//}
//
//func getAccessToken(authCode string) (*ZoomTokenResponse, error) {
//	// 构建获取访问令牌的请求
//	data := url.Values{}
//	data.Set("grant_type", "authorization_code")
//	data.Set("code", authCode)
//	data.Set("redirect_uri", redirectURI)
//
//	req, err := http.NewRequest("POST", tokenEndpoint, ioutil.NopCloser(strings.NewReader(data.Encode())))
//	if err != nil {
//		return nil, err
//	}
//
//	// 设置请求头
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	req.SetBasicAuth(clientID, clientSecret)
//
//	// 发送请求
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	// 解析响应
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var tokenResponse ZoomTokenResponse
//	if err := json.Unmarshal(body, &tokenResponse); err != nil {
//		return nil, err
//	}
//
//	return &tokenResponse, nil
//}
