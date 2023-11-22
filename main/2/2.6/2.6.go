package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// 定义一个结构体

	// 创建结构体实例
	//rec := Rectangle{Length: 10, Width: 20}
	rec := Rectangle{10, 20}

	area := rec.Area()
	fmt.Println(area)

	circle := Circle{10}
	fmt.Println(circle.Area())
}

// 1. 矩形结构体
type Rectangle struct {
	Length float64
	Width  float64
}

// 2. 矩形结构体的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// 3. 圆结构体
type Circle struct {
	Radius float64
}

// 4. 圆结构体的 Area 方法
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type T struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
	Address   struct {
		Street string `json:"street"`
		City   string `json:"city"`
		Zip    string `json:"zip"`
	} `json:"address"`
}
