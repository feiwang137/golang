/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 22:37
*/
package main

import (
	"fmt"
	"sync"
)

func main()  {
	 defer  fmt.Println("1")  // 在一个函数退出之前执行
	 defer  fmt.Println("2")
	 defer  fmt.Println("3")
	loopFunc()
}

func loopFunc()  {
	lock := sync.Mutex{}  // 创建一个lock 锁
	for i := 0;i<3 ;i++{
		go func() {
			lock.Lock() // 上锁
			defer lock.Unlock()  // 解锁
  			fmt.Println("loopFunc:",i)
		}()
	}
}