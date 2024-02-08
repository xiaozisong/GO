package main

// 多包导入
import (
	"fmt"
	"unsafe"
)

func main() {
	// 变量  数据类型
	// int8的范围是 -128 ~ 127
	var num1 int8 = 127
	fmt.Println(num1)

	// 无符号类型 uint8的范围是 0 ~ 255
	var num2 uint8 = 255
	fmt.Println(num2)

	var num3 byte = 18
	// 如何查看当前数据的类型 如果使用printf，后面必须用T%占位
	fmt.Printf("当前num3的类型是: T% \n", num3)

	// 如何查看当前数据的大小 存储字节
	var num4 int8 = 100
	fmt.Println(unsafe.Sizeof(num4))
}
