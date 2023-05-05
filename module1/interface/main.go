/*
@Description: xxxx
@Author: feiwang
@Date:  2023/5/5 14:40
*/

package main

import "fmt"

// 定义一个接口，这个接口里面有一个getName方法，只要实现了这个方法，就说明实现了这个接口
type IF interface {
	getName() string
}

type Human struct {
	firstName string
	lastName string
}

type plane struct {
	vendor string
	model string
}

// 方法 ； 格式：接收者 + 方法名 + 返回值（这部分和func就类似了）
func (p plane) getName() string{
	return fmt.Sprintf("vendor: %s , model: %s",p.vendor, p.model)
}

func (h Human) getName() string{
	return h.firstName + "," + h.lastName
}

func main()  {
	fmt.Println(">>>>>>>>> interface  <<<<<<<<<<<")
	interfaces := []IF{}

	h := new(Human)
	h.firstName="li"
	h.lastName="si"
	interfaces = append(interfaces, h)

	p := new(plane)
	p.vendor = "testVendor"
	p.model = "testModel"
	interfaces = append(interfaces, p)

	for _, f := range interfaces{
		fmt.Println(f.getName())
	}

	//
	fmt.Println(">>> interface test")
	var i IF
	newP := new(plane)
	newP.model = "newmodel"
	newP.vendor = "newVendor"

	i = newP
	returnResult := i.getName()
	fmt.Println(returnResult)
}

/*
剩余学习点：
1.反射 reflect
2.json编解码
3.错误处理

 */

