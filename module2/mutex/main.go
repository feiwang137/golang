/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/13 18:28
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {

	go lock()
	go rLock()
	go wLock()
	//go testDefer()
	time.Sleep(5 * time.Second)
}

func lock()  {
	// 互斥锁
	lock := sync.Mutex{}
	for i:=0;i<3;i++{
		lock.Lock() // 当第二次循环时，获取不到这个lock的，互斥。
		fmt.Println("lock ",i)
		defer lock.Unlock()  // 在子进程正式退出之前运行
	}

}

func rLock(){
	// 读写分离锁，不限制只读，限制锁并发读写和写
	lock := sync.RWMutex{}
	for i:=0;i<3;i++{
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rLock ",i)
	}
}


func wLock(){
	// 并发读写限制
	lock := sync.RWMutex{}
	for i:=0;i<3;i++{
		lock.Lock()
		defer  lock.Unlock()
		fmt.Println("wLock", i)
	}

}



func testDefer(){
	for i:=0 ; i<3; i++{
		//fmt.Println("i ",i)
		defer fmt.Println("done")
	}
	fmt.Println("end")
}
