package main

import (
)

type GreaterThan struct{
	Name string 
	Indexes []int
}

func (GreaterThan *GreaterThan) RunOnVariable(val ...interface{}) bool {
	f1, ok := getNum(val[0])
	if !ok {
		return false
	}
	f2, ok := getNum(val[1])
	if !ok {
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


type LessThan struct{
	Name string 
	Indexes []int
}

func (LessThan *LessThan) RunOnVariable(val ...interface{}) bool {
	f1, ok := getNum(val[0])
	if !ok {
		return false
	}
	f2, ok := getNum(val[1])
	if !ok {
		return false
	}

	if f1 < f2 {
		return true
	}
	return false
}

func (LessThan *LessThan) GetName() string {
	return LessThan.Name
}

func (LessThan * LessThan) GetArgIndexes() []int{
	return LessThan.Indexes
}


type EqualTo struct{
	Name string 
	Indexes []int
}

func (EqualTo *EqualTo) RunOnVariable(val ...interface{}) bool {
	f1, ok := getNum(val[0])
	if !ok {
		return false
	}
	f2, ok := getNum(val[1])
	if !ok {
		return false
	}

	if f1 == f2 {
		return true
	}
	return false
}

func (EqualTo *EqualTo) GetName() string {
	return EqualTo.Name
}

func (EqualTo * EqualTo) GetArgIndexes() []int{
	return EqualTo.Indexes
}