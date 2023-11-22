package main

import "fmt"

var (
	x = 1000000
	y = 20000
	z = "hello worlds"
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("hh")
	i := 10
	fmt.Println("i:", i)
	var x int
	fmt.Println(x)
	var y string
	fmt.Println(y)
	var z float32
	fmt.Println(z)
	var a = "string"
	fmt.Println(a)
	var var1, var2, var3 = 1, 2.9, 3
	fmt.Println(var1 + var3)
	fmt.Println(var2)

	var4, var6 := 1, 3
	fmt.Println(var4 + var6)

}
