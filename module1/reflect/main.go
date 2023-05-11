/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 11:24
*/
package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("reflect")
	myMap := make(map[string]string,10)
	myMap["a"] = "b"
	myMap["c"] = "d"

	// reflect
	// get type
	t:= reflect.TypeOf(myMap)
	fmt.Println(t)

	v := reflect.ValueOf(myMap)
	fmt.Println(v)

	//struct
	myStruct := T{A:"a",B:"b",C:"c"}

	v1 := reflect.ValueOf(myStruct)
	for i:=0; i < v1.NumField() ;i++{
		// 遍历map里面的值和index
		//fmt.Println(i,v1.Field(i))
		fmt.Printf("%d, %v\n",i,v1.Field(i))
	}

	// 看看这个struct实现了几个方法
	fmt.Println("to range map: numMethod")
	for i:=0; i< v1.NumMethod();i++{
		fmt.Printf("method %d: %v\n",i, v1.Method(i))
	}

	// call后面传入的nil表示方法需要接受的参数，nil表示传入一个空值
	result := v1.Method(0).Call(nil)
	fmt.Println(result)

}

type T struct {
	A string
	B string
	C string
}

// 基于struct实现一个方法； 注意这里是一个值copy
func (t T) String( ) string  {
	//fmt.Println(arg)
	return t.A + t.B + t.C +"1"
}