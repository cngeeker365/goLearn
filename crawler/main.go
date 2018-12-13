package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode  != http.StatusOK {
		fmt.Println("Error: status code is ", resp.StatusCode)
		return
	}


	//utf8Reader :=transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	encoding := determineEncoding(resp.Body)
	utf8Reader :=transform.NewReader(resp.Body,  encoding.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	fmt.Println()
}


func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ :=charset.DetermineEncoding(bytes, "")
	return e
}
