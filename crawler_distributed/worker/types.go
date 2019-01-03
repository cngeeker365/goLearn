package worker

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	"errors"
	"fmt"
	"log"
)

//{"ParseCityList", nil}, {"ProfileParser", userName}
type SerializedParser struct {
	FuncName	string
	Args 		interface{}
}

//engine中定义的部分无法在网络中传播，因此在这里构建一套类似的，同时需要能够进行转换
type Request struct {
	Url	string
	Parser SerializedParser
}

type ParseResult struct {
	Items 		[]engine.Item
	Requests 	[]Request
}

func SerializeRequest(r engine.Request) Request{
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			FuncName: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult{
	result := ParseResult{Items:r.Items}

	for _, req := range r.Requests{
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error){
	parser, err:=DescrializeParser(r.Parser)
	if err!=nil{
		return engine.Request{},nil
	}
	return engine.Request{
		Url: r.Url,
		Parser: parser,
	},nil
}

func DescrializeParser(p SerializedParser) (engine.Parser,error) {
	switch p.FuncName {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity),nil
	case config.NilParser:
		return engine.NilParser{},nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok{
			return parser.NewProfileParser(userName),nil
		}else{
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}

	default:
		return nil, errors.New("unknow parser name")
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult{
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _,req := range r.Requests{
		engineReq, err:=DeserializeRequest(req)
		if err!=nil{
			log.Printf("error deserializing request: %v\n", err)
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
