/*
@Description: xxxx
@Author: feiwang
@Date:  2023/5/4 16:37
*/

package main

import "fmt"

const LuckyNum = 37

func main()  {
	fmt.Println("goland study function")
	// 接受外部的参数，参数解析
	//fmt.Println(os.Args)
	//name := flag.String("name", "value","your name")
	//flag.Parse()
	//fmt.Println(*name)

	DoOperation(1, increase)

}

func init()  {
	// init function
	var newSlice = []int{}
	name := "lisi"
	//var myLuckyNum = 0

	fmt.Printf("%T\n", LuckyNum)
	fmt.Println(name," lucky number is ",LuckyNum)
	newNewSlice := append(newSlice, 1,2,3,4,5)
	fmt.Println(newNewSlice)

}


// callback function
func increase(a,b int){
	println("increase result is:",a+b)
}

func decrease(a,b int){
	println("decrease result is:", a-b)
}

func DoOperation(y int, f func(int, int)){
	f(y,1)
}



