package main

import (
)

type Between struct {
	Name    string
	Min     float64
	Max     float64
	Outside bool
	Indexes []int
}

func (Between *Between) RunOnVariables(val ...interface{}) bool {
	f, ok := getNum(val[0])
	if !ok {
		return false
	}

	if (Between.Min <= f && Between.Max >= f) != Between.Outside {
		return true
	}
	return false
}

func (Between *Between) GetName() string {
	return Between.Name
}
func (Between *Between) GetArgIndexes() []int {
	return Between.Indexes
}

type Minimum struct {
	Name    string
	Min     float64
	Indexes []int
}

func (Minimum *Minimum) RunOnVariables(val ...interface{}) bool {
	f, ok := getNum(val[0])
	if !ok {
		return false
	}

	if Minimum.Min <= f {
		return true
	}
	return false
}

func (Minimum *Minimum) GetName() string {
	return Minimum.Name
}
func (Minimum *Minimum) GetArgIndexes() []int {
	return Minimum.Indexes
}

type Maximum struct {
	Name    string
	Max     float64
	Indexes []int
}

func (Maximum *Maximum) RunOnVariables(val ...interface{}) bool {
	f, ok := getNum(val[0])
	if !ok {
		return false
	}

	if Maximum.Max >= f {
		return true
	}
	return false
}

func (Maximum *Maximum) GetName() string {
	return Maximum.Name
}
func (Maximum *Maximum) GetArgIndexes() []int {
	return Maximum.Indexes
}
