/*
@Description: xxxx
@Author: feiwang
@Date: 2023/4/24 21:30
*/


package main

import "fmt"

//func sup_egg()  {
//	fmt.Println("this is a suprise egg.")
//
//}




func main()  {

	//var num int
	//num = 6

	// if
	//fmt.Println("study if ")

	//if num == 5 {
	//	fmt.Printf("%d > 5 \n" ,num)
	//} else if  num == 10 {
	//	fmt.Println("num > 10")
	//} else if num == 15 {
	//	fmt.Println("num > 15")
	//} else {
	//	fmt.Println("what's fuck...")
	//}
	
	// switch if的进化版本
	//fmt.Println("=== study switch")
	//switch num  {
	//	case 2: fmt.Println("is 2")
	//	case 4: fmt.Println("is 4")
	//case 6:
	//	fmt.Println("skip")
	//	sup_egg()
	//case 8: fmt.Println("is 8")
	//default:
	//	fmt.Println("what's the fuck.")


	// 计数器循环: 初始化语句; 条件语句; 修饰语句
	//for i:=0; i<10; i++{
	//	fmt.Println(i)
	//}

	// 类似于while 循环
	//sum := 995
	//for ; sum < 1000; {
	//	//sum += sum
	//	sum +=1
	//	fmt.Println(sum)
	//}

	//变量和常量
	//var name string
	//name = "wangfei"
	//fmt.Println(name)

	//const name2 = "lisi"
	//fmt.Println(name2)

	//var name1,name2 string
	//name1 ,name2 = "zhangsan","lisi"
	//fmt.Println(name1,name2)
	//
	//var result bool
	//result = false
	//

	//var name string = "wangfei"
	//
	//// 类型推导
	//name2 := name
	//fmt.Println(name2)

	// 显式转换类型
	//var num int
	//num = 12
	//var numFloat float64
	//numFloat = float64(num)
	//fmt.Printf("%f",numFloat)

	// 数组 指定了长度
	//var testArry [5] int   // 这种写法就是容量和长度都是5
	////fmt.Printf("1 %v %v\n",len(testArry), cap(testArry))
	//testArry = [5]int{1,2,3,4,5}
	//fmt.Println(testArry[4])
	////fmt.Printf("2 %v %v",len(testArry), cap(testArry))
	////fmt.Printf("2 %v",len(testArry))
	//
	//mySliceTest := testArry[1:4]
	//fmt.Println(mySliceTest)
	//fullSlice := testArry[:] // 整个arry转换为了slice
	//fmt.Println(fullSlice)
	////删除index 为2的值 ；需要单独写一个func
	////fullSlice
	//remove3item := deleteItem(fullSlice,2)
	//fmt.Println(remove3item)

	// 切片，指定长度
	//var testSlice []int
	//fmt.Println(len(testSlice))
	//testSlice = []int{1,2,3,4,5}
	//fmt.Println(len(testSlice))
	//fmt.Println(testSlice)

	// make 和 new ；
	// new出来的slice 没有初始化，只分配了内存地址。

	//mySlice1 := new([]int)
	//fmt.Println(& mySlice1)
	//
	//sliceTest := []int{1,23}
	//
	//mySlice1 = & sliceTest
	//fmt.Println(& mySlice1)

	// make出来slice 初始化了
	//mySlice2 := make([]int,2,2)
	//
	//*mySlice1 = mySlice2
	//fmt.Println(& mySlice1)
	//
	//fmt.Println(mySlice2, len(mySlice2),cap(mySlice2))
	//mySlice2 = []int{2,3}
	//fmt.Println(mySlice2, len(mySlice2),cap(mySlice2))
	//
	//mySlice3 := make([]int,0)
	//fmt.Println(len(mySlice3))
	//mySlice3 = []int{1,2}
	//fmt.Println(len(mySlice3))

	//aSlice := []int{1,2,3}
	//bSlice := []int{}
	//
	//cSlice := aSlice
	//
	//fmt.Println(aSlice)
	//fmt.Println(bSlice)
	//fmt.Println(cSlice)
	//
	//fmt.Println("------")
	//
	//aSlice = append(aSlice, 2)  // 此时a切片的内存地址发生改变，aslice 将不等于 cslice
	//fmt.Println(aSlice)
	//fmt.Println(bSlice)
	//fmt.Println(cSlice)
	//aSlice[2] = 66
	//fmt.Println(aSlice)
	//
	//for _,v := range aSlice{
	//	fmt.Printf("value is : %v\n", v)
	//}


	//var dict map[string]string
	//dict := make(map[string]string)
	//
	//dict["name"] = "wangfei"
	//dict["addr"] = "anhui"
	//
	//value, exists := dict["name"]
	//if exists{
	//	fmt.Println(value)
	//
	//}
	//fmt.Println(exists, value)
	//
	//for k,v := range dict{
	//	fmt.Println(k,v)
	//}


	//myFuncMap := map[string]func() int{"funcA": func() int{return  1}}
	//
	//
	//
	//
	//f := myFuncMap["funcA"]
	//fmt.Println(f())
	//
	//
	//funcMap := map[string]func() int{} // 这个int表示返回值的类型，不懂 先过吧。

	//myStructfunc()
	action1()

}

type myType struct {
	Name string
}

func myStructfunc() {

	st := myType{Name:"fuck"}

	printMyType(&st)
}

func printMyType(t *myType){
	fmt.Println(t.Name)
}

func  deleteItem(slice []int, index int) []int  {
	// 这是一个删除指定index的func
	return append(slice[:index], slice[index+1:]...) // 这里的...不知道什么意思。
}

func action1() {
	fmt.Println("hello world")
	taskArray := [5]string{}
	taskArray = [5]string{"I","am","stupid","and","weak"}
	fmt.Println(taskArray)

	for index,value := range taskArray{
		fmt.Println(index,value)
		if value == "stupid"{
			taskArray[index] = "smart"
		} else if value == "weak"{
			taskArray[index] = "strong"
		}
	}
	fmt.Println(taskArray)

	const fuck = 22
	fmt.Println(fuck)

	type ageType int

	type ServiceType string
	const (
		ServiceTypeClusterIP ServiceType = "ClusterIP"
		ServiceTypeNodePort ServiceType = "NodePort"
		wangfeiAge ageType = 31
		age2 int = 18
	)

	var num ageType
	num = 13

	fmt.Println(ServiceTypeClusterIP)
	fmt.Println(wangfeiAge,num, age2)


}



