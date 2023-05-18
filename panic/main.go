/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 22:44
*/
package main

import "fmt"

func main()  {
	fmtPrintMsg()
	fmt.Println("I Come Back.")  // 这个证明了 从panic中恢复过来了，如果灭有recover() ，不会执行这段代码的。
}

func fmtPrintMsg()  {
	defer func() {  // 在函数最后执行
		fmt.Println("defer func is called")
		if err:= recover(); err != nil{  // 从错误中回复
			fmt.Println(err)
		} else {
			//fmt.Println("continue...")
		}
	}()
	//panic("a panic is triggered")  // 直接crash当前的进程，后面的代码就不会执行了。
	fmt.Println("egg...")

}



