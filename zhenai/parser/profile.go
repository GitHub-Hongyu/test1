package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/model"
)
//<div class="m-btn purple" data-v-bff6f798>45岁</div>
//<div class="m-btn purple" data-v-bff6f798>168cm</div>
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-312fdcc4>([\d^<]+)</h1>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	//	<h1 class="nickName" data-v-312fdcc4>为谁结婚</h1>
	profile := model.Profile{}

	profile.Name = name
	//爬年龄
	age, _ := strconv.Atoi(extractString(contents, ageRe))
	profile.Age = age

	//爬身高
	height, _ := strconv.Atoi(extractString(contents, heightRe))

	profile.Height = height

	//爬姓名等等
	//profile.Name = extractString(contents, nameRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
