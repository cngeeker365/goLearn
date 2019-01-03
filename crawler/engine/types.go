package engine

type ParserFunc func(contents []byte, url string) ParseResult
type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize()	(name string, args interface{})
}

type Request struct {
	Url string
	//ParserFunc func([]byte) ParseResult
	//ParserFunc ParserFunc
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items	 []Item
}

type Item struct {
	Url		string
	Type 	string
	Id 		string
	Payload	interface{}
}
//---------------------------------------------------------------
type NilParser struct {}

func (NilParser) Parse(_ []byte, _ string) ParseResult{
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}){
	return "", nil
}
//---------------------------------------------------------------
type FuncParser struct {
	Parser ParserFunc
	Name string
}

func (f *FuncParser) Parse(content []byte, url string) ParseResult {
	return f.Parser(content, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

//工厂函数，用于构建Parser
func NewFuncParser(p ParserFunc, name string) *FuncParser{
	return &FuncParser{
		Parser: p,
		Name: name,
	}
}



//func NilParser([]byte) ParseResult{
//	return ParseResult{}
//}