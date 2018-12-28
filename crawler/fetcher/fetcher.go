package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string)([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode  != http.StatusOK {
		fmt.Println("Error: status code is ", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}


	//utf8Reader :=transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)
	utf8Reader :=transform.NewReader(bodyReader,  encoding.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("%s\n", all)
	//fmt.Println()
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