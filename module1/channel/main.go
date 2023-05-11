/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/7 16:31
*/
package main

import (
	"fmt"
	"math/rand"

	//"math/rand"

	//"reflect"
	"time"
)

func main()  {
	ch := make(chan int,10)

	go func() {
		fmt.Println("Hello World")
		ch <- 2  // 向channel传送数据
	}()
	i := <- ch  // 从channel读取数据
	fmt.Println(i)


	//defer  close(ch)
	//if v,notClose := <- ch; notClose{  // notClose 的值是 True / False
	//	fmt.Println(notClose,v)
	//}

	//channelBuffer()
	//channelDemo()
	//signChannel()
}

func channelDemo(){
	for i:=0;i<5;i++{
		go fmt.Println(i)  // 启动一个新的协程
	}
	time.Sleep(time.Second)

}

func channelBuffer(){
	ch := make(chan int ,2)
	//defer close(ch)
	//
	go func() {
		for i:=0; i<10; i++{
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(8)
			ch <- n
		}
		defer close(ch)
	}()

	fmt.Println("hello from main")
	for v:= range ch{
		fmt.Println("receiving: ", v)
	}
}

func signChannel()  {
	fmt.Println("sign channel")

	ch := make(chan int)
	go prod(ch)
	go consume(ch)
	time.Sleep(time.Second)
	fmt.Println("ffff")
}

func prod(ch chan <- int)  {
	//for {ch <- 1}
	ch <- 1
	//defer close(ch)
}

func consume(ch <- chan int ) {
	//for {
	//	v := <- ch
	//	fmt.Println(v)
	//}

	//for {
		v := <- ch
		fmt.Println(v)
	//}
}


// 其它
/*
1. select : 可以通过此去检查不同的通道，通道是否有数据传过来。
  // 随机选择检查
	select {
	case v:= <- ch :
		fmt.Println("to do something")
	case v:= <- ch2:
		fmt.Println("to do something")
	default:
		fmt.Println("to do something")
	}

2. timer : 定时器，配合select使用。

	timer := time.NewTimer(time.Second)
	select {
		// check normal channel
		case <-ch:
			fmt.Println("received from ch")
		case <-timer.C:
			fmt.Println("timeout waiting from channel ch")
	}

3. close

	ch := make(chan int)
	defer close(ch)
	if v, notClosed := <-ch; notClosed {
		fmt.Println(v)
	}

4。 context 高级用法简化

 */