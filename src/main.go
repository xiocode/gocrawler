/**
 * Created with IntelliJ IDEA.
 * User: xio
 * Date: 13-2-18
 * Time: 下午4:21
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"gocrawler"
	"fmt"
)

func main() {
	fmt.Printf("Hello world!")
	crawler := &gocrawler.Crawler{}
	crawler.GET("http://www.baidu.com")
}
