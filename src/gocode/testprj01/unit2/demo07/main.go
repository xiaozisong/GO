package main

import "fmt"

func main() {
	// 基本数据类型的转换，在go语言中，是不允许进行隐式转换的，只可以显式转换(强转)
	// 转换的方式为 T(v) T(Type): 类型 V(Value): 值
	var a1 int8 = 11
	// var a2 int64 = a1; // 这行编译不会通过，直接报错
	var a3 int16 = int16(a1)

	fmt.Println(a3)

	// 如果是从高位转低位会造成溢出的情况
	var a4 int32 = 11111
	var a5 int8 = int8(a4)

	fmt.Println(a5)

	// var a6 int32 = a1 + 12 // 编译不通过，因为需要显示转换a1

	var a7 int8 = int8(a4) + 127 // 编译通过，但是会出现溢出的情况
	// var a8 int8 = int8(a4) + 128 // 编译不通过，因为128已经超出了int8的表示范围
	fmt.Println(a7)
	// fmt.Println(a8)
}
