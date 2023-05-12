/*
@Description: xxxx
@Author: feiwang
@Date:  2023/5/11 16:37
*/

/*
作业内容：
• 基于Channel编写一个简单的单线程生产者消费者模型

• 队列:
   队列长度10，队列元素类型为 int

• 生产者:
   每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞

• 消费者:
   每1秒从队列中获取一个元素并打印，队列为空时消费者阻塞.
*/


package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	channelLength := 10
	ch := make(chan int, channelLength)

	// producer
	go func() {
		for {
			// print now channel length
			if  len(ch)== channelLength{
				fmt.Println(time.Now().Format("2006-05-02 15:04:05"), "full, channel blocked.")
			} else {
				fmt.Println(time.Now().Format("2006-05-02 15:04:05")," channel length: ", len(ch))
			}

			// send number to ch
			rand.Seed(time.Now().UnixNano())
			i:=rand.Intn(1000)
			//fmt.Println("send i>: ",i)
			ch <- i
			i++
			time.Sleep(1 * time.Second)
		}

	}()

	// consumer
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C{
		select {
		case <- ch:
			fmt.Println(time.Now().Format("2006-05-02 15:04:05")," recv i< : ", <- ch)
		default:
			fmt.Println(time.Now().Format("2006-05-02 15:04:05")," waiting for channel:ch to send data.")
		}
	}
}


