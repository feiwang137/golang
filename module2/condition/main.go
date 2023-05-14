/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/14 13:42
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	 q := Queue{
	 	queue: []string{},
	 	cond: sync.NewCond(&sync.Mutex{}),
	 }

	 go func() {
	 	for {
	 		q.Enqueue("a")
	 		time.Sleep(2 * time.Second)
		}
	 }()

	 for {
	 	result:=q.Dequeue()
	 	fmt.Println(result)
	 	time.Sleep(2 * time.Second)
	 }

}


type Queue struct {
	queue []string
	cond *sync.Cond  // sync.cond ； 让goroutine在满足特定的情况下被唤醒执行。
}


func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.queue) == 0 {
		fmt.Println("no data avaiable, wait.")
		q.cond.Wait()
	}

	result := q.queue[0]
	q.queue = q.queue[:1]

	return result
}

func (q *Queue) Enqueue(item string)   {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	//fmt.Printf("putting #{item} to queue, notify all\n")
	fmt.Printf("putting %s to queue, notify all\n", item)
	q.cond.Broadcast()
}