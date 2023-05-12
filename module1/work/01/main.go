/*
@Description: xxxx
@Author: feiwang
@Date:  2023/5/12 13:29
*/

package main

import "fmt"

func main(){
	fmt.Println("hello world")
	//taskArray := [5]string{}
	taskArray := [5]string{"I","am","stupid","and","weak"}
	fmt.Println(taskArray)

	for index,value := range taskArray{
		//fmt.Println(index,value)
		if value == "stupid"{
			taskArray[index] = "smart"
		} else if value == "weak"{
			taskArray[index] = "strong"
		}
	}
	fmt.Println(taskArray)

}
