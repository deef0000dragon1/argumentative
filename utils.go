package main

import (
	"errors"
)
func getString(in interface{}) (string, bool) {

	s,ok := in.(string)


	if !ok {
		return "", false
	}
	return s, true
}

func getBool(in interface{}) bool, bool{
switch(in.type) {
	case string:
	case bool:
		return in.(bool), true
}
return false, false
}

func getInt64(in interface{}) int64, bool{

}

func getInt32(in interface{}) int32, bool{

}

func getInt16(in interface{}) int16, bool{

}

func getInt8(in interface{}) int8, bool{

}

func getUInt64(in interface{}) uint64, bool{

}

func getUInt32(in interface{}) uint32, bool{

}

func getUInt16(in interface{}) uint16, bool{

}

func getUInt8(in interface{}) uint8, bool{

}

func getInt(in interface{}) int, bool{

}

func getByte(in interface{}) byte, bool{

}

func getRune(in interface{}) rune, bool{

}



func getFloat64(in interface{}) float64, bool{

}

func getFloat32(in interface{}) float64, bool{

}


func getComplex64(in interface{}) complex64, bool{

}
func getComplex128(in interface{}) complex128, bool{

}



func getNum(in interface{}) float64, bool{
	return getFloat64(in)
}