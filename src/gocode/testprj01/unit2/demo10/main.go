package main

import "fmt"

func main() {
	// 指针，指针就是当前变量所存储值的内存地址，指针指向这个内存地址，也可以拿到该内存中的值
	var v1 int = 18
	// 定义指针类型使用 *int, 取内存地址的值使用&，如果想调用出指针所指向内存中的值，使用*取值
	var point *int = &v1
	fmt.Println(point)
	fmt.Println(*point)
}
