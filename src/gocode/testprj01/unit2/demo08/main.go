package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 其他数据类型转换为字符串
	// 方式1重点使用，方式2使用次数较少，因为过于繁琐
	// 有两种方式 1. 使用fmt.sprintf() 2. strconv.Formatxxx
	var n1 int8 = 10
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Println(s1)

	var n2 float64 = 1.11
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Println(s2)

	// %t针对于布尔值
	var b1 bool = true
	var s3 string = fmt.Sprintf("%t", b1)
	fmt.Println(s3)

	// %c针对字符类型
	var c1 byte = '1'
	var s4 string = fmt.Sprintf("%c", c1)
	fmt.Println(s4)

	// 2. strconv
	var n3 int = 8
	// formatint方法的第一个参数只接收int64类型的数据，所以需要先进行转换，第二参数表示进制，10就是10进制，2就是2进制
	var s5 string = strconv.FormatInt(int64(n3), 10)
	fmt.Println(s5)

	var n4 float32 = 2.22
	// formatFloat方法的第一个参数只接收float64，第二个参数是格式化的类型，f代表数据的表示方法，f:ddd.ddd(小数的十进制表示)，
	// 除了f还有b、e、E、g、G
	// 第三个参数表示保留9位小数，第四个参数代表float64类型
	var s6 string = strconv.FormatFloat(float64(n4), 'f', 9, 64)
	fmt.Println(s6)

	var b2 bool = false
	var s7 string = strconv.FormatBool(b2)
	fmt.Println(s7)
}
