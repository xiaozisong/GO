package main

import "fmt"

func main() {
	// 获取用户终端输入，有两种方式
	// 1. 使用scanln
	// 输入学生的年龄、姓名、成绩、是否为VIP
	var age int
	// 输入的时候使用指针的方式去影响变量
	fmt.Println("请输入年龄")
	// fmt.Scanln(&age)

	var name string
	fmt.Println("请输入姓名")
	// fmt.Scanln(&name)

	var score float32
	fmt.Println("请输入成绩")
	// fmt.Scanln(&score)

	var isVIP bool
	fmt.Println("是否为VIP")
	// fmt.Scanln(&isVIP)

	// fmt.Printf("学生的年龄: %v, 学生的姓名: %v, 学生的成绩: %v, 是否为VIP: %v", age, name, score, isVIP)

	// 2. 使用Scanf，每次中断使用空格
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Printf("学生的年龄: %v, 学生的姓名: %v, 学生的成绩: %v, 是否为VIP: %v", age, name, score, isVIP)

}
