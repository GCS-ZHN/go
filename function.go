package main

import "fmt"

const PI float64 = 3.141562654

func main() {
	fmt.Printf("max(5, 2) = %d\n", max(5, 2))
	var a, b string = swap("Google", "Runoob")
	fmt.Println(a, b)
	var c Circle
	c.radius = 2
	fmt.Printf("半径为%.2f的圆面积为%.2f\n", c.radius, c.getArea())
}

// go的函数定义类似python，返回值类型放在后面

// 返回a与b的最大值
func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 和go一样，返回值可以是多个
func swap(x string, y string) (string, string) {
	return y, x
}

type Circle struct {
	radius float64
}

// go并没有直接面向对象，通过给函数指定所有者，实现面向对象封装，这种函数称为方法method
func (this Circle) getArea() float64 {
	return PI * this.radius * this.radius
}
