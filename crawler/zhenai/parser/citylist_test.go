package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err :=fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err:= ioutil.ReadFile("citylist_test_data.html")
	if err!=nil {
		panic(err)
	}
	result :=ParseCityList(contents)

	t.Logf("result have %d requests\n", len(result.Requests))
	t.Logf("result have %d items\n", len(result.Items))

}
