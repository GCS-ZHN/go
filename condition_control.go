package main

import "fmt"

const ad int = 10

// go中的左花括号不允许单独成行，这或许是开发者的编程风格吧
func main() {
	var test_name string
	fmt.Println("请输入测试类型")
	fmt.Scan(&test_name)
	switch test_name {
	case "if":
		if_test()
	case "switch":
		switch_test()
	case "for":
		for_test()
	case "while":
		while_test()
	case "foreach":
		foreach_test()
	default:
		fmt.Println("Not implemented test type: " + test_name)
	}

}

func if_test() {
	var age int
	fmt.Println("请输入年龄")
	fmt.Scan(&age)
	var name string = "郑奇奇"
	fmt.Printf("Hello, %s, ago is %d\n", name, age)
	if age < 100 {
		fmt.Println("傻乎乎")
	} else if age < 200 {
		fmt.Println("哈士奇")
	} else {
		fmt.Println("哈哈哈")
	}
}

func switch_test() {
	var grade string
	fmt.Println("请输入成绩等级字母")
	fmt.Scan(&grade)
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
}

// go支持经典的for循环形式
func for_test() {
	var times int
	fmt.Println("请输入循环次数")
	fmt.Scan(&times)
	for i := 0; i < times; i++ {
		fmt.Printf("%d,", i)
	}
}

// go支持经典的while循环形式，但都用for关键字
func while_test() {
	var times, i int
	fmt.Println("请输入循环次数")
	fmt.Scan(&times)
	for i < times {
		fmt.Printf("%d,", i)
		i++
	}
}

// go支持经典的foreach循环，与python不同的是，用for, range组合而非for , in
// 支持 for k := range data 和 for k, v := range data
func foreach_test() {
	var data = [5]int{1, 2, 5, 6, 7}
	for idx := range data {
		fmt.Printf("%d,", data[idx])
	}
}
