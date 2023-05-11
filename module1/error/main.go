/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 22:24
*/
package main

import (
	"errors"
	"fmt"
)

func main()  {
	err := fmt.Errorf("this is an error")
	err2 := errors.New("this is error")
	fmt.Println(err)
	fmt.Println(err2)
	err = errorTest()
	if err != nil{
		fmt.Println(err)
	}

}

func errorTest() error{
	num := 11
	if num > 10{
		return nil
	}else {
		err2 := errors.New("this is error")
		//msg := "i'm ok"
		return err2
	}

}


// k8s 通过interface 实现 error报错信息的归类。

