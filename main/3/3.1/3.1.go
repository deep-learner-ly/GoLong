package main

import (
	"fmt"
	"math"
)

// 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func main() {
	circle := Circle{3}
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())

	dog := Dog{"狗狗"}
	dog.Speak()
	dog.Eat()

	cat := Cat{"猫猫"}
	cat.Speak()
	cat.Eat()
}

type Animal interface {
	Speak() string
	Eat() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return d.name + " " + "汪汪"
}

func (d Dog) Eat() string {
	return d.name + " " + "啃骨头"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return c.name + " " + "喵喵"
}

func (c Cat) Eat() string {
	return c.name + " " + "吃鱼"
}

//src = source
//srv = server
//arg = argument
//conn = connect, connection
//attr = attribute
//abs = absolute
//min = minimum
//len = length
//auth = authenticate
//buf = buffer
//ctl = control
//ctx = context
//str = string
//msg = message
//fmt = format
//dest = destination
//diff = difference
//orig = original
//recv = receive
//ref = reference
//repo = repository
//util = utility
//fmt = format
