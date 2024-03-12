package main

import "fmt"

func main() {
	// 浮点数可能会出现精度丢失问题，如想使精度更加准确建议使用float64
	var num1 float32 = 3.14

	var num2 float32 = -3.14

	var num3 float32 = 314e+2

	var num4 float64 = 314e+2

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
