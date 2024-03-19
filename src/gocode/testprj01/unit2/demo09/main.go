package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将string转换为基本数据类型 -- 使用strconv.parsexxx
	// parseint 与 parsefloat的bitSize参数不管你传什么，最后都会输出64类型的数据，但是它可以转换为32或者是8、16类型而不改变它的值。
	var s1 string = "11"
	var n1 int64
	// parsexxx方法返回的值有两个 一个value，一个error，需要用两个变量去接，如果不想接，可以使用_去忽略
	// parseint方法第一个参数是要转换的字符串，第二个参数是进制，第三个参数是类型 int8 int16 32 64..
	n1, _ = strconv.ParseInt(s1, 10, 8)
	fmt.Println(n1)

	var s2 string = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(s2)
	fmt.Println(b1)

	var s3 string = "3.33"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Println(f1)

	// 如果转换时string的值与基本类型不对应会发生什么呢
	var s4 string = "golang"
	var n2 int64
	n2, _ = strconv.ParseInt(s4, 10, 8) // 默认为0，因为golang不是一个有效的int类型数据
	fmt.Println(n2)

	var b2 bool
	b2, _ = strconv.ParseBool(s4)
	fmt.Println(b2) // 默认为false

	var f2 float64
	f2, _ = strconv.ParseFloat(s4, 64)
	fmt.Println(f2)
}
