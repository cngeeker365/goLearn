package fetcher

import (
	"awesomeProject/crawler_distributed/config"
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//http.Get不能过快，可能会出发目标网站的反扒机制
//突破反扒的机制：如限速、更换UserAgent、虚拟IP

var rateLimiter = time.Tick(time.Second/config.Qps)

func Fetch(url string)([]byte, error) {
	<- rateLimiter

	log.Printf("Fetching url %s", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode  != http.StatusOK {
		fmt.Println("Error: status code is ", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)
	utf8Reader :=transform.NewReader(bodyReader,  encoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v \n", err)
		return unicode.UTF8
	}
	e, _, _ :=charset.DetermineEncoding(bytes, "")
	return e
}