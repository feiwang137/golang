/*
@Description: xxxx
@Author: feiwang
@Date:  2023/5/10 10:55
*/

package main

import (
	"fmt"
	"time"
)

//func main()  {
//	fmt.Println("prod consumer")
//
//	ch := make(chan int ,10)
//
//	// this prod; send data in channel
//	go func(){
//		for i:=0; i<8; i++{
//			ch <- i
//			fmt.Println("i is: ",i)
//
//		}
//		fmt.Println("send done..")
//		//defer close(ch)
//	}()
//
//
//	go func() {
//		for i:=1; i<cap(ch);i++{
//			time.Sleep(time.Second)
//
//			fmt.Println("[consumer1] recv i: ", <- ch)
//		}
//	}()
//
//	go func() {
//		for i:=1;i< cap(ch);i++{
//			time.Sleep(time.Second)
//			fmt.Println("[consumer2] recv i: ", <- ch)
//		}
//	}()
//
//	// 如果没有close channel的话，下面的消费者就会hang住，然后达到20s后，就会退出进程。
//	time.Sleep(20 * time.Second)
//
//	fmt.Println("finish")
//}

// 遍历通道缓冲区
//func  main()  {
//	// 必须设置一个随机的种子，系统会默认使用一个固定的种子，否则rand每次生成的随机数都是一样的。
//	//rand.Seed(time.Now().UnixNano())
//	//func(){
//	//	for i:=0;i<10;i++{
//	//		n := rand.Intn(1000)
//	//		fmt.Println(n)
//	//		time.Sleep(time.Second)
//	//	}
//	//}()
//
//	rand.Seed(time.Now().UnixNano())
//	ch := make(chan int,10)
//	go func() {
//		defer close(ch)
//
//		for i:=0;i<10;i++{
//			n := rand.Intn(10)
//			ch <- n
//		}
//	}()
//
//	for v:=range ch {
//		fmt.Println(v)
//	}
//	//time.Sleep(10*time.Second)
//
//}

// 单向通道
//func main()  {
//	c := make(chan  int)
//	go prod(c)
//	go consume(c)
//	close(c)
//
//	//time.Sleep(2 * time.Second)
//	if v,notClosed := <- c; notClosed{
//		fmt.Println(notClosed)
//		fmt.Println("v is: ",v)
//	} else {
//		fmt.Println(notClosed)
//		fmt.Println("channel is closed.")
//	}
//
//}
//
//func  prod(ch chan <- int)  {
//	//for {ch<-1}
//	ch <- 1
//	ch <- 2
//}
//
//func consume(ch <- chan int ){
//	fmt.Println(<- ch)
//	//for {<-ch}
//	//Invalid operation: ch <- 2 (send to receive-only type <-chan int)
//	//ch <- 2
//}
//

//func main(){
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	go func(){
//		for i:=0 ; i< 10 ; i++{
//			ch1 <- i
//			time.Sleep(1* time.Second)
//		}
//		defer close(ch1)
//	}()
//
//	go func(){
//		for i:=0 ; i< 10 ; i++{
//			ch2 <- i
//			time.Sleep(2* time.Second)
//		}
//		defer close(ch2)
//	}()
//
//	for {
//		// 通过select去轮询channel（随机），如果通道阻塞就走default
//		select {
//		case v:= <- ch1:
//			fmt.Println("ch1 v: ",v)
//		case v:= <- ch2:
//			fmt.Println("ch2 v: ",v)
//		default:
//			fmt.Println("waiting...")
//		}
//		time.Sleep(1 * time.Second)
//	}
//
//}

// 节拍器 结合 channel完成一个生产者消费者模型
func main()  {
	ch1 := make(chan int,10)
	done := make(chan bool)
	defer close(ch1)

	// consumer
	go func(){
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C{
			select {
				case <- done:
					fmt.Println("child process is interrupt...")
					return
				default:
					fmt.Printf("send messages:%d \n", <-ch1)
					//fmt.Println(<-done)
			}
		}
	}()


	// producer
	for i:=0;i<10;i++{
		ch1 <- i
		//time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Second) // 这里等待5s是为了，让consumer消费
	close(done)  // 关闭done channel，select会捕捉到done关闭了。
	time.Sleep(1 * time.Second)  // 这里决赛点，如果CPU先调度运行了主进程 main process 将退出，
	fmt.Println("main process exit!")

}

