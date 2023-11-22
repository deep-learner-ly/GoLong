package main

import "fmt"

func main() {
	// 声明数组
	var arr [5]int

	// 初始化数组
	arr = [5]int{1, 2, 3, 4, 5}

	arr1 := [5]int{1, 2, 3, 4, 5}

	// 访问数组元素
	fmt.Println(arr[0]) // 输出: 1

	fmt.Println(arr1[2])

	// 声明切片
	var slice []int

	// 使用make函数创建切片
	slice = make([]int, 3, 6)

	// 初始化切片

	fmt.Println("切片长度:", len(slice)) // 输出: 3
	fmt.Println("切片容量:", cap(slice)) // 输出: 5
	fmt.Println("切片内容:", slice)      // 输出: [0 0 0]
	slice = []int{1, 2, 3}

	fmt.Println("切片长度:", len(slice)) // 输出: 3
	fmt.Println("切片容量:", cap(slice)) // 输出: 5
	fmt.Println("切片内容:", slice)

	// 使用 make 创建一个字符串到整数的映射
	m := make(map[string]int)

	// 添加键值对
	m["one"] = 1
	m["two"] = 2

	fmt.Println("映射内容:", m) // 输出: map[one:1 two:2]

	// 创建一个动态数组
	list := []int{1, 2, 3, 4, 5}

	sublist := list[1:3]
	fmt.Println(sublist)

	// 创建一个二维切片
	matrix := [][]int{{1, 2, 3}, {8, 4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix[1][3])

}
