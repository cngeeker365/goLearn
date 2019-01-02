package controller

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/frontend/model"
	"awesomeProject/crawler/frontend/view"
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler{
	client, err:=elastic.NewClient(elastic.SetSniff(false))
	if err !=nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client:client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	q := strings.TrimSpace(req.FormValue("q"))

	from, err:= strconv.Atoi(req.FormValue("form"))
	if err!=nil{
		from=0
	}

	page, err := h.getSearchResult(q, from)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))). //支持将搜索条件 Age:(<30) 默认修改为 Payload.Age:(<30)
		From(from).
		Do(context.Background())
	if err !=nil{
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PreFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string{
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	re.ReplaceAllString(q, "Payload.$1:")
}

