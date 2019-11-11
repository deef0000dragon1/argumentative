package main

import (
	"fmt"
	"strconv"
)
func getString( in interface {})(string, bool) {

	fmt.Print("")
	s, ok := in .(string)


	if !ok {
		return "", false
	}
	return s, true
}

func getBool( in interface {})(bool, bool) {

	switch in.(type) {
		case string:
			b, err := strconv.ParseBool(in.(string))
			if err != nil {
				return false, false
			}
			return b, true
		case bool:
			return in.(bool), true
	}
	return false, false
}

func getNum( in interface {})(float64, bool) {
	return getFloat64( in )
}


func getInt64( in interface {})(int64, bool) {

			switch in.(type) {
		case string:
			i, err := strconv.ParseInt(in.(string), 0,64)
			if err != nil {
				return 0, false
			}
			return i, true
		case int64:	return in.(int64), true
		case int32:	return int64(in.(int32)), true
		case int16:	return int64(in.(int16)), true
		case int8:	return int64(in.(int8)), true
		case uint64:return int64(in.(uint64)), true
		case uint32:return int64(in.(uint32)), true
		case uint16:return int64(in.(uint16)), true
		case uint8:	return int64(in.(uint8)), true //also byte
		case int:	return int64(in.(int)), true

	}
	return 0, false
}
func getFloat64( in interface {})(float64, bool) {
		switch in.(type) {
		case string:
			f, err := strconv.ParseFloat(in.(string), 64)
			if err != nil {
				return 0, false
			}
			return f, true
		case float64:
			return in.(float64), true
		case float32:
			return 0,true
	}
	return 0, false
}



//TODO LATER
func getInt32( in interface {})(int32, bool) {
	return 0, false
}

func getInt16( in interface {})(int16, bool) {
	return 0, false
}

func getInt8( in interface {})(int8, bool) {
	return 0, false
}

func getUInt64( in interface {})(uint64, bool) {
	return 0, false
}

func getUInt32( in interface {})(uint32, bool) {
	return 0, false
}

func getUInt16( in interface {})(uint16, bool) {
	return 0, false
}

func getUInt8( in interface {})(uint8, bool) {
	return 0, false
}

func getInt( in interface {})(int, bool) {
	i, b := getInt64( in )
	return int(i), b
}

func getByte( in interface {})(byte, bool) {
	return getUInt8(in)
}

func getRune( in interface {})(rune, bool) {
	return 0, false
}




func getFloat32( in interface {})(float32, bool) {
	return 0, false
}


func getComplex64( in interface {})(complex64, bool) {
	return 0, false
}
func getComplex128( in interface {})(complex128, bool) {
	return 0, false
}