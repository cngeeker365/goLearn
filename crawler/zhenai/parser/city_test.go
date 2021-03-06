package parser

import (
	"awesomeProject/crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCity(t *testing.T) {
	const url = "http://m.zhenai.com/zhenghun/aba"
	contents, err :=fetcher.Fetch(url)
	//contents, err:= ioutil.ReadFile("citylist_test_data.html")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
	result :=ParseCity(contents,url)

	t.Logf("result have %d requests\n", len(result.Requests))
	t.Logf("result have %d items\n", len(result.Items))
	for _,r := range result.Items{
		fmt.Println(r)
	}
	for _,q := range  result.Requests{
		fmt.Println(q.Url)
	}

}
