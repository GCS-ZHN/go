package main

import "fmt"

// 全局变量，在函数外声明，不一定要被实际使用
var a int = 10

func main() {
	/*
	 局部变量，作用域在函数内
	 函数返回值、函数参数以及函数内声明的变量都属于局部变量
	 其中声明变量必须被使用，否则会编译失败，提示 declared but not used
	*/
	var b int = 2
	fmt.Printf("Global variable a = %d\n", a)
	fmt.Printf("Local variable b = %d\n", b)

	/*
		虽然使用var来定义变量，但和js的var不一样，go变量是基于代码块的作用域
		而js则只有全局变量和函数内变量两种，与python一样。个人认为块作用域
		更加合理。
	*/
	{
		var c int = 3
		fmt.Printf("Local variable c = %d\n", c)
	}

	// 函数参数不一定要被使用，这其实是给接口实现以最大自由
	test(1, 2)
}

func test(a, b int) {
	fmt.Printf("No parameter were used, which was legal")
}
