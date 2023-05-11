/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 12:40
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	//strHuman := "{Name:wangfei,Age:18}"
	//newStruct := unmarshal2Struct(strHuman)
	//fmt.Println("----")
	//fmt.Println(newStruct)

	myStruct := Human{Name: "wangFei",Age: 31}
	strStruct := marshal2JsonString(myStruct)
	fmt.Println(strStruct)
	fmt.Printf("%T",strStruct)

	fmt.Println("----")
	newStruct := unmarshal2Struct(strStruct)
	fmt.Println(newStruct)
	unmarshalAnyJson(strStruct)


}

type Human struct {
	Name string
	Age int
}

// string to struct
func unmarshal2Struct(humanStr string) Human  {
	h := Human{}
	err := json.Unmarshal([]byte(humanStr),&h)
	if err != nil{
		fmt.Println(err)
	}
	return h
}

// struct to string
func marshal2JsonString(h Human) string{
	//h.Age = 30
	//h.Name = "liSi"
	updatedBytes,err := json.Marshal(&h)
	if err !=nil{
		fmt.Println(err)
	}
	return string(updatedBytes)
}


// any json to struct
func unmarshalAnyJson(humanStr string) {
	var obj interface{}
	err := json.Unmarshal([]byte(humanStr), &obj)
	if err != nil{
		fmt.Println(err)
	}
	// 这里的obj是将动态类型转换为 map[string]interface{} 的类型
	objMap, _ := obj.(map[string]interface{})  // 使用了类型断言，将 obj 转换为 map[string]interface{} 类型。
	fmt.Printf("%T\n",objMap)
	r := objMap["Age"]
	fmt.Println("=======",r,"====")
	fmt.Println(objMap)
	for k,v := range objMap{
		switch  value:= v.(type) {  // v.(type)类型断言的语法，只能用于interface{}类型的变量
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}
}