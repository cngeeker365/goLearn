package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
	"strings"
)

//const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityRegex = `<a href="(http://m.zhenai.com/u/[0-9]+)#seo" class="left-item" data-v-4062b8db><div class="right-item" data-v-4062b8db><div class="f-nickname" data-v-4062b8db>([^<]+)<span class="u-gender0" data-v-4062b8db></span>`

func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityRegex)
	matches := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	
	for _,m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+strings.TrimSpace(string(m[2])))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	return result
}