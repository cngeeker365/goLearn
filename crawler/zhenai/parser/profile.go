package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
	"strconv"
)

var ageReg = regexp.MustCompile(``)
var marriageReg = regexp.MustCompile(``)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err != nil {
		//user age is age
		profile.Age = age
	}

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) > 2 {
		return string(match[1])
	}
	return ""
}
