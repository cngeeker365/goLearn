package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityListReg  = `{linkContent:"([^"]+)",linkURL:"(http://m.zhenai.com/zhenghun/[0-9a-z]+)"}`

func ParseCityList(contents []byte) engine.ParseResult{
	re:=regexp.MustCompile(cityListReg)
	matches :=re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}

	limit := 10
	for _,m:= range matches {
		result.Items = append(result.Items, "City "+string(m[1]))
		result.Requests = append(	result.Requests,
									engine.Request{
										Url: string(m[2]),
										ParserFunc: ParseCity,
									})

		limit--
		if limit==0{
			break
		}
	}
	return result
}
