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

	for _,m:= range matches {
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(	result.Requests,
									engine.Request{
										Url: string(m[2]),
										ParserFunc: engine.NilParser,
									})
		//fmt.Printf("City: %s,\t url: %s\n", m[1], m[2])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
