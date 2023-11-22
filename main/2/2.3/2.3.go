package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x大于5")
	} else {
		fmt.Println("x不大于5")
	}

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("星期一")
	case "Tuesday":
		fmt.Println("星期二")
	default:
		fmt.Println("其他")
	}

	// 基本的for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 类似于while的循环
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	num := 7

	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	switchTest()

	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

}

func switchTest() {
	var num1, num2 int
	var operator string

	fmt.Print("输入第一个数字: ")
	fmt.Scanln(&num1)

	fmt.Print("输入第二个数字: ")
	fmt.Scanln(&num2)

	fmt.Print("输入运算符(+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("无效的运算符")
	}
}
