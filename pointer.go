package main

import "fmt"

func main() {
	/* go的指针声明方式与C类似，使用*来声明
	在类型前面加*即为指针的声明
	未初始化的指针值为nil，即空指针，等价于0
	*/
	var a *int
	fmt.Println(a)
	fmt.Printf("空指针nil格式等价于%x\n", a)

	// 和C一样，使用&来获取变量的指针
	var a_value = 10
	a = &a_value
	fmt.Println(a)
	fmt.Printf("Point value is a hex value address %x\n", a)

	// 和C一样，使用*来对指针进行解引用
	fmt.Printf("指针a所指地址%x储存的值为%d\n", a, *a)

	/* 指针可以形成指针数组，元素是指针
	 */
	b := [3]int{10, 100, 200}
	var ptr [3]*int
	for i := 0; i < len(b); i++ {
		fmt.Printf("数据%d的地址%x\n", i, &b[i])
		ptr[i] = &b[i]
	}
	fmt.Printf("%x\n", &ptr)

	/* go和C语言一样，不是直接的OOP语言，
	其参数传递均为按值传递，即形参不影响实参
	若需要按引用传递，则需要使用指针类型的参数
	所谓高级语言的引用概念，实际上还是基于指针
	*/
	var d = 10
	selfAdd(&d)
	fmt.Println(d)
}

func selfAdd(value *int) {
	(*value)++
}
