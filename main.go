package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/manticoresoftware/go-sdk/manticore"
	"github.com/segmentio/ksuid"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	ID        string `json:"ent_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"created_at"`
}
type ManticoreDocRequest struct {
	Index string `json:"index,omitempty"`
	ID    int64  `json:"id,omitempty"`
	Doc   Data   `json:"doc,omitempty"`
}
type ManticoreResponse struct {
	Index   string `json:"_index"`
	ID      int64  `json:"_id"`
	Created bool   `json:"created"`
	Result  string `json:"result"`
	Status  int    `json:"status"`
}
type ManticoreSearchRequest struct {
	Index  string `json:"index"`
	Source struct {
		Includes []string `json:"includes"`
	} `json:"_source"`
	Query struct {
		QueryString string      `json:"query_string"`
		Bool        interface{} `json:"bool"`
	} `json:"query"`
	Highlight struct {
		Fields        []string `json:"fields"`
		Limit         int      `json:"limit"`
		LimitSnippets int      `json:"limit_snippets"`
		BeforeMatch   string   `json:"before_match"`
		AfterMatch    string   `json:"after_match"`
	} `json:"highlight"`
	Limit int `json:"limit"`
}

func getSearchRequest() ManticoreSearchRequest {
	return ManticoreSearchRequest{
		Index: "test",
		Source: struct {
			Includes []string `json:"includes"`
		}{
			Includes: []string{
				"ent_id",
				"title",
				"content",
				"created_at"},
		},
		Query: struct {
			QueryString string      `json:"query_string"`
			Bool        interface{} `json:"bool"`
		}{
			QueryString: "",
		},
		Highlight: struct {
			Fields        []string `json:"fields"`
			Limit         int      `json:"limit"`
			LimitSnippets int      `json:"limit_snippets"`
			BeforeMatch   string   `json:"before_match"`
			AfterMatch    string   `json:"after_match"`
		}{
			Fields: []string{
				"title",
				"content",
			},
			Limit:         40,
			LimitSnippets: 5,
			BeforeMatch:   `<em&nbsp;id=%SNIPPET_ID%>`,
			AfterMatch:    "</em>",
		},
		Limit: 10,
	}
}

type IdMustPredicate struct {
	Must []struct {
		Equals struct {
			EntId string `json:"title"`
		} `json:"equals"`
	} `json:"must"`
}

func GetIdMustPredicate(id string) IdMustPredicate {
	return IdMustPredicate{
		Must: []struct {
			Equals struct {
				EntId string `json:"title"`
			} `json:"equals"`
		}{
			{
				Equals: struct {
					EntId string `json:"title"`
				}{
					EntId: id,
				},
			},
		},
	}
}

type ManticoreSearchResponse struct {
	Took     int           `json:"took"`
	TimedOut bool          `json:"timed_out"`
	Hits     ManticoreHits `json:"hits"`
}
type ManticoreHits struct {
	Total int                  `json:"total"`
	Hits  []ManticoreDocOutput `json:"hits"`
}
type ManticoreDocOutput struct {
	//ID        string                `json:"_id"`
	Score     int                   `json:"_score"`
	Source    Data                  `json:"_source"`
	Highlight ManticoreDocHighlight `json:"highlight"`
}
type ManticoreDocHighlight struct {
	Title   []string `json:"title,omitempty"`
	Content []string `json:"content,omitempty"`
	//AuthorName []string `json:"authorName,omitempty"`
	//CreatedAt  []string `json:"created_at,omitempty"`
	//Source     []string `json:"source,omitempty"`
}

func main() {
	ctx := context.Background()
	cl := manticore.NewClient()
	cl.SetServer(os.Getenv("MANTICORE_HOST"), 9312)
	cl.SetConnectTimeout(3 * time.Second)
	_, err := cl.Open()
	if err != nil {
		log.Fatalln("connect to manticore failed", err)
	}
	ress, err := cl.Sphinxql("create table test(ent_id string, title string, content text, created_at text) charset_table='0..9, chinese' morphology='icu_chinese' stopwords='zh' html_strip='1'")
	fmt.Println(ress, err)

	id := ksuid.New().String()
	req := ManticoreDocRequest{
		Index: "test",
		Doc: Data{
			ID:        id,
			Title:     "Title",
			Content:   "Content",
			CreatedAt: time.Now().Format("2006年1月2日 15:04:05"),
		}}
	b, _ := json.Marshal(req)
	resp, err := http.Post("http://localhost:9308/insert", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println("request manticore failed", err)
	}
	defer resp.Body.Close()
	var manticoreRes ManticoreResponse
	_ = json.NewDecoder(resp.Body).Decode(&manticoreRes)
	log.Println(manticoreRes.ID)

	searchReq := getSearchRequest()
	searchReq.Query.QueryString = "title"
	searchReq.Query.Bool = GetIdMustPredicate("title")
	b, _ = json.Marshal(searchReq)
	reqq, _ := http.NewRequest(http.MethodPost, "http://localhost:9308/search", bytes.NewReader(b))
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	reqq = reqq.WithContext(ctx)
	client := http.DefaultClient
	resp, err = client.Do(reqq)
	log.Println(resp.StatusCode)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err)
	}
	defer resp.Body.Close()
	var res ManticoreSearchResponse
	_ = json.NewDecoder(resp.Body).Decode(&res)
	log.Printf("搜索字段：%s, 命中数量：%+v", "title", res.Hits.Total)
	log.Println(res.Hits.Hits)
	for _, v := range res.Hits.Hits {
		log.Println(v)
	}
}
