package main

import "fmt"

func main() {
	// 指针的四个注意细节
	// 1.指针指向的变量是可变的
	var num int = 10
	var point *int = &num
	fmt.Println(num)
	*point = 20
	fmt.Println(num)

	// 2.指针变量接收的一定是地址值
	// var num2 int = 11
	// var point2 *int = 11 // 这样是不被允许的

	// 3.指针变量的地址不可以不匹配
	// var num3 int = 22
	// var point3 *float32 = &num3 // float类型指针只能指向float类型的数据，int对应int以此类推

}
