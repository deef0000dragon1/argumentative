package main

import (
	"strconv"
)

type GreaterThan struct{
	Name string 
	Indexes []int
}

func (GreaterThan *GreaterThan) RunOnVariable(val ...interface{}) bool {
	str1, ok := val[0].(string)
	if !ok {
		return false
	}

	f1, err := strconv.ParseFloat(str1, 64)
	if err != nil {
		return false
	}
	str2, ok := val[1].(string)
	if !ok {
		return false
	}

	f2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		return false
	}
	if f1 > f2 {
		return true
	}
	return false
}

func (GreaterThan *GreaterThan) GetName() string {
	return GreaterThan.Name
}

func (GreaterThan * GreaterThan) GetArgIndexes() []int{
	return GreaterThan.Indexes
}