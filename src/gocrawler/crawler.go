package gocrawler

import (
	"net/http"
	. "gocrawler/utils"
	"fmt"
)

type Crawler struct {
	URL        string
	Headers    map[string]string
	Resp       http.Response
	HttpClient http.Client
}

type result struct {
	URL     string
	Content string
	Length  uint
}

func (crawl *Crawler) SetHeaders(headers map[string]string) {
	crawl.Headers = headers
}

func (crawl *Crawler) GET(url string) {
	if url != "" {
		crawl.URL = url
	}
	resp, err := crawl.HttpClient.Get(crawl.URL)
	CheckErr(err)
	defer resp.Body.Close()
	fmt.Println(resp)
}


