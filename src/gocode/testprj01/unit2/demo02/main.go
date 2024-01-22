package main

import "fmt"

// 声明全局变量
var globalVar1 int = 666
var globalVar2 int = 777

// 以上多次声明很麻烦，设计者提供简易的声明方式
var (
	number1        = 20
	number2        = 30
	age2    string = "tom"
)

func main() {

	// 声明在{}中的变量称作局部变量

	// 定义变量的几种形式
	// 1. 直接声明并赋值
	var name = "xiaoming"
	// 2. 声明类型并赋值
	var age int = 18
	// 3. 简易写法，省略var，使用:=的写法
	height := 188
	// 4. 声明变量，并定义类型，但不赋值，系统会使用默认值
	var weight float32
	fmt.Println(name, age, height, weight)

	// go中也支持一次性声明多变量
	var n1, n2, n3 = 10, 20, 30
	fmt.Println(n1, n2, n3)

	var n4, n5, n6 = 10, false, "jack"
	fmt.Println(n4, n5, n6)

	var n7, n8, n9 int = 20, 30, 40
	fmt.Println(n7, n8, n9)
}
