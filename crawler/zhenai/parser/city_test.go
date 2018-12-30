package parser

import (
	"awesomeProject/crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err :=fetcher.Fetch("http://m.zhenai.com/zhenghun/aba")
	//contents, err:= ioutil.ReadFile("citylist_test_data.html")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
	result :=ParseCity(contents)

	t.Logf("result have %d requests\n", len(result.Requests))
	t.Logf("result have %d items\n", len(result.Items))
	for _,r := range result.Items{
		fmt.Println(r)
	}
	for _,q := range  result.Requests{
		fmt.Println(q.Url)
	}

}
