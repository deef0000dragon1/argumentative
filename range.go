package main

import (
	"strconv"
)

type Between struct {
	Name    string
	Min     float64
	Max     float64
	Outside bool
	Indexes []int
}

func (Between *Between) RunOnVariable(val ...interface{}) bool {
	str, ok := val[0].(string)
	if !ok {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
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

func (Minimum *Minimum) RunOnVariable(val ...interface{}) bool {
	str, ok := val[0].(string)
	if !ok {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
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

func (Maximum *Maximum) RunOnVariable(val ...interface{}) bool {
	str, ok := val[0].(string)
	if !ok {
		return false
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
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
