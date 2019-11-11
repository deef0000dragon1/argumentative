package main

import "strings"

type Agreement struct {
	Validations []Validation //rename to facts
	validators  map[string]bool
	Truth       *Truths
}

type Validation interface {
	GetName() string
	GetArgIndexes() []int
	RunOnVariables(...interface{}) bool
}

type Truths struct {
	Validation string
	Signifier  string
	Inputs     []*Truths
}

func (Agreement *Agreement) Argue(arguments ...interface{}) bool {

	validators := make(map[string]bool)
	for _, v := range Agreement.Validations {
		validators[v.GetName()] = validate(v, arguments)
	}
	return Agreement.close(*Agreement.Truth)
}

func validate(val Validation, arguments ...interface{}) bool {
	inputs := make([]interface{}, 0)

	for _, v := range val.GetArgIndexes() {
		inputs = append(inputs, arguments[v])
	}
	return val.RunOnVariables(inputs)
}

func (arg Agreement) close(t Truths) bool {

	m := arg.validators
	if t.Validation != "" {
		r, ok := m[t.Validation]
		if !ok || !r {
			return false
		}
		return true
	}

	switch strings.ToLower(t.Signifier) {

	case "and":
		for _,v := range t.Inputs {
			out := arg.close(*v)
			if !out {
				return false
			}
		}
		return true

	case "or":
		for _,v := range t.Inputs {
			out := arg.close(*v)
			if out {
				return true
			}
		}
		return false
		
	case "not":
		if len(t.Inputs) < 1{
			return true
		}
		return arg.close(*t.Inputs[0])
	}
	return false
}
