package main

import (
	"fmt"
	"sync"
	"net/http"
	"io/ioutil"
	. "gocrawler/utils"
)

type Fetcher interface {
	Fetch(url string) (result fetchResult)
}

func Crawl(url string, fetcher Fetcher, out chan string, end chan bool) {
	fmt.Println(url)

	if _, ok := crawled[url]; ok {
		end <- true
		return
	}

	result := fetcher.Fetch(url)
	body, url, err := result.body, result.url, result.err
	if err != nil {
		out <- fmt.Sprintln(err)
		end <- true
		return
	}

	out <- fmt.Sprintf("found: %s %q\n", url, body)
	end <- true
	fmt.Println("END! %s", url)

	crawledMutex.Lock()    //上个锁
	crawled[url] = true
	crawledMutex.Unlock()  //解锁
}

var crawled = make(map[string]bool)
var crawledMutex sync.Mutex

//锁

func main() {
	out := make(chan string)
	end := make(chan bool)

	urls := []string{"http://www.baidu.com", "http://www.qq.com"}
	for _, url := range urls {
		go Crawl(url, fetcher, out, end)
	}

	var result interface {}

//	for i := 0; i < len(urls); i++ {
//		<-end
//	}

	for {
		select {
		case result = <-out:
			fmt.Println(result)
		case result = <-end:
			if len(crawled) == len(urls) {
				fmt.Println("Finished!")
				return
			}
		}
	}

}

type fetchResult struct {
	body     string
	url      string
	err      error
}

func (crawl *Crawler) Fetch(url string) (result fetchResult) {
	if url, ok := url, crawl.ok; ok {
		if url != "" {
			crawl.URL = url
		}
		resp, err := crawl.HttpClient.Get(crawl.URL)
		CheckErr(err)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		CheckErr(err)
		if resp.StatusCode == 200 {
			body := string(bodyBytes)
			result := fetchResult{body:body, url:url, err:err}
			return result
		}
	}
	return fetchResult{body:"", url:"", err:fmt.Errorf("错误")}
}

type Crawler struct {
	URL        string
	ok         bool
	Headers    map[string]string
	Resp       http.Response
	HttpClient http.Client
}

func (crawl *Crawler) SetHeaders(headers map[string]string) {
	crawl.Headers = headers
}

var fetcher = &Crawler{ok:true}
