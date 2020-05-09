package main

import (
	"fmt"
	"reflect"
)

func main() {

	var xInt = 2
	fmt.Println("Type:",reflect.TypeOf(xInt).String())
	fmt.Println("Value:",xInt)

	var xInt8 = int8(2)
	fmt.Println(reflect.TypeOf(xInt8).String())
	fmt.Println("Value:",xInt8)

	var xInt16 = int16(2)
	fmt.Println(reflect.TypeOf(xInt16).String())
	fmt.Println("Value:",xInt16)

	var xInt32= int32(2)
	fmt.Println(reflect.TypeOf(xInt32).String())
	fmt.Println("Value:",xInt32)

	var xInt64= int64(2)
	fmt.Println(reflect.TypeOf(xInt64).String())
	fmt.Println("Value:",xInt64)

	var xFloat64 = 2.
	fmt.Println(reflect.TypeOf(xFloat64).String())
	fmt.Println("Value:",xFloat64)

	var xFloat32 = float32(2.)
	fmt.Println(reflect.TypeOf(xFloat32).String())
	fmt.Println("Value:",xFloat32)

	var data interface{} // Type switch allows switching on the type of something instead of value
	data = ""
	data = int64(2)
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Println(c, "is an int64\n")
	}

}
