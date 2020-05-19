package parser

import (
	"engine"
	"fmt"
	"model"
	"regexp"
	"strconv"
)

var (
	AgeRe        = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
	HeightRe     = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d+]cm)</div>`)
	WeightRe     = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d+]kg)</div>`)
	IncomeRe     = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">月收入:([^<]*)</div>`)
	Marriage     = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(未婚)</div>`)
	EducationRe  = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(大学本科)</div>`)
	OccupationRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(.*师)</div>`)
	HokouRe      = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">籍贯:([^<]*)</div>`)
	XinzouRe     = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(.*座\([\d]+.[\d]+-[\d]+.[\d]+\))</div>`)
	WorkLocRe    = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">工作地:([^<]*)</div>`)
)

func ParseProfile(content []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Age = extractInt(content, AgeRe)
	profile.Height = extractString(content, HeightRe)
	profile.Income = extractString(content, IncomeRe)
	profile.Marriage = extractString(content, Marriage)
	profile.Occupation = extractString(content, OccupationRe)
	profile.Education = extractString(content, EducationRe)
	profile.Hokou = extractString(content, HokouRe)
	profile.Xinzuo = extractString(content, XinzouRe)
	profile.WorkLoc = extractString(content, WorkLocRe)
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	fmt.Println(profile)
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(content []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		age, err := strconv.Atoi(string(match[1]))
		// user age is age
		if err != nil {
			return 0
		}
		return age
	}
	return 0
}
