package model

type SearchResult struct {
	Query 		string
	PreFrom		int
	NextFrom	int
	Hits		int64
	Start 		int
	Items 		[]interface{}
}
