package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	APPID          = "wx1caa725933424656"
	APPSECRET      = "6476ac5584649e76533af9bc6c3a8525"
	SentTemplateID = "tPJaHWWj1vzh9BhNO-37m36ibUk3pcft7xuEjDJFmOU" //每日一句的模板ID，替换成自己的
	WeatherVersion = "v1"
)

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type sentence struct {
	Content     string `json:"content"`
	Note        string `json:"note"`
	Translation string `json:"translation"`
}

func main() {
	everydaysen()
}

//发送每日一句
func everydaysen() {

	t := time.Now().Format("2006-01-02 15:04:05")
	title := "XX银行员工考核对网点绩效的影响"
	fxurl := "https://www.baidu.com"
	access_token := getaccesstoken()
	if access_token == "" {
		return
	}

	flist := getflist(access_token)
	if flist == nil {
		return
	}
	reqdata := "{\"keyword1\":{\"value\":\"" + t + "\"}, \"keyword2\":{\"value\":\"" + title + "\"}, \"keyword3\":{\"value\":\"无\"}, \"keyword4\":{\"value\":\"测试\"}, \"keyword5\":{\"value\":\"2次\"}, \"remark\":{\"value\":\"1111111111111111\"}}"
	//reqdata := "{\"content\":{\"value\":\"" + aa + "\", \"color\":\"#0000CD\"}, \"note\":{\"value\":\"" + aa + "\"}, \"translation\":{\"value\":\"" + aa + "\"}}"
	for _, v := range flist {
		msg(access_token, reqdata, fxurl, SentTemplateID, v.Str)
	}
}

//获取微信accesstoken
func getaccesstoken() string {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", APPID, APPSECRET)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取微信token失败", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("微信token读取失败", err)
		return ""
	}

	token := token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("微信token解析json失败", err)
		return ""
	}

	return token.AccessToken
}

//获取关注者列表
func getflist(access_token string) []gjson.Result {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + access_token + "&next_openid="
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取关注列表失败", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败", err)
		return nil
	}
	flist := gjson.Get(string(body), "data.openid").Array()
	return flist
}

func msg(access_token string, reqdata string, fxurl string, templateid string, openid string) {

	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token

	reqbody := "{\"touser\":\"" + openid + "\", \"template_id\":\"" + templateid + "\", \"url\":\"" + fxurl + "\", \"data\": " + reqdata + "}"

	PostHeader(url, []byte(reqbody))
	//log.Println(url)
	//log.Println(reqbody)
}

type Responses struct {
	Message string `json:"message"`
}

// PostHeader 发送post请求
func PostHeader(url string, msg []byte) ([]byte, error) {
	client := &http.Client{}
	var resData Responses
	req, err := http.NewRequest("POST", url, bytes.NewReader(msg))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		//解析返回数据
		json.NewDecoder(resp.Body).Decode(&resData)
		return nil, status.Error(codes.Internal, "PostRequest failed "+resData.Message)
	}
	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodys, nil
}
