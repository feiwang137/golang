/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/13 19:06
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	fmt.Println("wait group")
	//waitBySleep()
	//waitByChannel()
	waitByWG()
}



func waitBySleep(){
	// 方法1 指定时间，等待程序执行完成。有问题时间小，子线程没有执行完成，时间多程序的效率就很低了。
	for i:=0;i<3;i++{
		go fmt.Println(i)
	}
	time.Sleep(5 * time.Second)
}

func waitByChannel(){
	// 方法2  通过channel的方式 生产者，消费者模型，让子进程正常执行完，再结束主进程。

	ch := make(chan bool,100)
	go func(){
		for i:=0; i<100 ;i++ {
			fmt.Println(i)
			ch <- true
		}
	}()

	for i:=0 ;i<100;i++{
		<- ch
	}

}


func waitByWG(){
	wg := sync.WaitGroup{}
	wg.Add(100) // 初始化，塞入100个线程
	//go func(){
	//	for i:=0 ; i< 100 ;i++{
	//		fmt.Println(i)
	//		wg.Done()    // 每个线程完成之后，执行Done 标记这个线程执行完成了。
	//	}
	//}()

	for i:=0 ; i<100; i++{
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)  // send i to Anonymous func.

	}

	// 等待所有的线程Done，执行完成
	wg.Wait()

}
