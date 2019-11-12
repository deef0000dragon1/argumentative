package main

import (
	"regexp"
)

type RegexMatch struct{
	Name string 
	Regex string
	Indexes []int
}

func (RegexMatch *RegexMatch) RunOnVariables(val ...interface{}) bool {
	str, ok := getString(val[0])
	if !ok {
		return false
	}

	re := regexp.MustCompile(RegexMatch.Regex)
	return re.Match([]byte(str))
	
}

func (RegexMatch *RegexMatch) GetName() string {
	return RegexMatch.Name
}
func (RegexMatch * RegexMatch) GetArgIndexes() []int{
	return RegexMatch.Indexes
}