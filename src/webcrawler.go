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

func Crawl(url string, fetcher Fetcher, out chan string, wg sync.WaitGroup) {
	fmt.Println(url)

	if _, ok := crawled[url]; ok {
		wg.Done()
		return
	}

	crawledMutex.Lock()    //上个锁
	crawled[url] = true
	crawledMutex.Unlock()  //解锁


	result := fetcher.Fetch(url)
	body, url, err := result.body, result.url, result.err
	if err != nil {
		out <- fmt.Sprintln(err)
		wg.Done()
		return
	}



	out <- fmt.Sprintf("found: %s %q\n", url, body)
	fmt.Println("DEBUG!")
	wg.Done()
	fmt.Println("END! %s", url)

}

var crawled = make(map[string]bool)
var crawledMutex sync.Mutex
var wg sync.WaitGroup

//锁

func main() {
	out := make(chan string)

	urls := []string{"http://www.baidu.com", "http://www.qq.com"}
	for _, url := range urls {
		wg.Add(1)
		go Crawl(url, fetcher, out, wg)
	}

	var result interface {}

	//	for i := 0; i < len(urls); i++ {
	//		<-end
	//	}

	wg.Wait()

	for {
		select {
		case result = <-out:
			fmt.Println(result)
		default:
			fmt.Println("Finished!")

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
