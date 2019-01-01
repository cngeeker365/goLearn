package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"regexp"
)

//var base = regexp.MustCompile(`<span data-v-10352ec0>([^<]+)</span>`)
var common = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([^<]+)</div>`)
var num = regexp.MustCompile(`([0-9]+)`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	//profile.Name = name
	//profile.Name = string(base.FindAllSubmatch(contents, 1)[0][1])


	commonMatches:=common.FindAllSubmatch(contents, -1)
	extractMatches(commonMatches, &profile)

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}

	profile.Info = append(profile.Info, name)

	return result
}

func extractMatches(matches [][][]byte, profile *model.Profile) {
	var result []string
	for _, m := range matches{
		result = append(result, string(m[1]))
	}
	profile.Info = result
	//profile.Marriage		= result[0]
	//profile.Age,_			= strconv.Atoi(num.FindString(result[1]))
	//profile.Xingzuo			= result[2]
	//profile.Height,_		= strconv.Atoi(num.FindString(result[3]))
	//profile.Weight,_		= strconv.Atoi(num.FindString(result[4]))
	//profile.WorkLocation 	= strings.Split(result[5],":")[1]
	//profile.Income 			= strings.Split(result[6],":")[1]
	//profile.Work			= result[7]
	//profile.Education		= result[8]
	//profile.Nation			= result[9]
	//profile.HuKou			= strings.Split(result[10],":")[1]
	//profile.Shape			= strings.Split(result[11],":")[1]
	//profile.Smoke			= result[12]
	//profile.Drink			= result[13]
	//profile.House			= result[14]
	//profile.Car				= result[15]
	//profile.HasChild		= result[16]
	//profile.WantChild		= strings.Split(result[17],":")[1]
}

