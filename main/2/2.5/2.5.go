package main

import (
	"fmt"
	"sort"
)

// 带有参数和返回值的函数
func add(a, b int) int {
	return a + b
}

func main() {
	//// 调用带有参数和返回值的函数
	//result := add(3, 5)
	//fmt.Println("3 + 5 =", result)

	// 匿名函数
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	// 闭包
	multiplier := func(factor int) func(x bool) int {
		return func(x bool) int {
			if x {
				return factor
			} else {
				return -factor
			}

		}
	}
	// 使用闭包创建乘法函数
	multiplyByTwo := multiplier(2)
	fmt.Println("Multiply by 2:", multiplyByTwo(false)) // 输出 8
	//闭包的用途：
	//封装状态： 闭包可以用于封装一些状态，使得函数可以访问和修改这些状态。
	//函数工厂： 通过闭包生成并返回其他函数。
	//回调函数： 在异步编程中，可以使用闭包作为回调函数来处理异步操作的结果。

	arr := []int{1, 4, 2, 1, 9}
	fmt.Println(findMax(arr))

	res := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", res)

	fmt.Println(mRet())
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func mRet() (int, string) {
	return 1, "hh"
}
