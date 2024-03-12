package main

import "fmt"

func main() {
	// 在go语言中是没有字符串这个概念的，底层会将字符转换成对应的ASCII，中文则是转换成unicode
	// byte类型数据可以直接参加运算
	var c1 byte = 'a'

	var c2 byte = '('

	var c3 byte = '6'

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	// 直接输出汉字会报错
	// var c4 byte = '中'
	// fmt.Println(c4)

	// 但是使用int来输出就不会报错
	var c5 int = '中'
	fmt.Println(c5)

	// 如果想打印字符串，需要使用到格式化输出中的c%
	var c6 byte = 'A'
	fmt.Printf("c5的对应字符串是: %c", c6)
}
