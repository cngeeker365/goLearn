package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error){
	//log.Printf("Fetching %s \n", r.Url)
	body, err:= fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher: error fetching url %s: %v \n", r.Url, err)
		return ParseResult{}, nil
	}
	return r.ParserFunc(body, r.Url), nil
}