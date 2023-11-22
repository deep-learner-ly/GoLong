package main

import "fmt"

func main() {
	// 声明一个整数变量
	var num1 int
	num1 = 10

	// 声明并初始化一个字符串变量
	var str1 string = "Hello, Go!"

	// 类型推断，声明时不指定类型
	var num2 = 5

	// 简短声明（仅限于函数内部），不使用 var 关键字
	name := "John"

	fmt.Println(num1, str1, num2, name)
	typeTest()
	printIntegerOperations(3, 4)

	printFloatComparison(3.0, 3.00)

	printStringConcatenation("he", "hhh")
	printBooleanOperations(true, false)
	myFunc(1, 1.5)
}

func typeTest() {
	// 整数类型
	var num1 int = 42
	var num2 int8 = 8
	var num3 uint = 20

	// 浮点数类型
	var float1 float32 = 3.14
	var float2 float64 = 2.718

	// 字符串类型
	var str1 string = "Hello"
	var str2 string = "Go"

	// 布尔类型
	var isTrue bool = true

	fmt.Println(num1, num2, num3, float1, float2, str1, str2, isTrue)
}

func printIntegerOperations(num1, num2 int) {
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	remainder := num1 % num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}
func printFloatComparison(float1, float2 float64) {
	isEqual := float1 == float2
	fmt.Println("Are the floating-point numbers equal?", isEqual)
}

func printStringConcatenation(str1, str2 string) {
	result := str1 + " " + str2
	fmt.Println("Concatenated String:", result)
}
func printBooleanOperations(isTrue, isFalse bool) {
	andResult := isTrue && isFalse
	orResult := isTrue || isFalse
	notResult := !isTrue

	fmt.Println("Logical AND:", andResult)
	fmt.Println("Logical OR:", orResult)
	fmt.Println("Logical NOT:", notResult)
}

func myFunc(x int, y float32) {
	fmt.Println("x:", x)
	fmt.Println("y", y)
}
