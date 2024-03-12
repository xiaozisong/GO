package main

import "fmt"

func main() {
	// 字符串类型
	var s1 string = "abcdefg"

	fmt.Println(s1)

	// 字符串类型的数据一但声明，就不可以改变其中的字符
	var s2 string = "abc"
	s2 = "def" // 不是说不可以替换，而是字符串其中的某一项字符值不可变
	// s2[0] = "e"
	fmt.Println(s2)

	// 字符串拼接
	var s3 string = "AAA" + "BBB"
	s3 += "CCC"

	fmt.Println(s3)

	// 如果字符串中有特殊字符，想省事，不想做转义拼接等，可以使用反引号 `` 与js类似
	var s4 = `
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

	`
	fmt.Println(s4)
}
