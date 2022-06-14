package main

import "fmt"

func main() {
	// 声明数组形式与变量类似，区别在于加了一个维度数据
	var data [10]int
	fmt.Println(data)

	/* 与变量类似，可以直接声明并初始化数组，
	与cpp/java的数组类似，在后面接“{}”来赋值
	赋值元素个数不超过限定的个数，不足的元素按默认
	值补齐
	*/
	var data1 = [5]float32{10.0, 2.1, 3.14, 7.0}
	fmt.Println(data1)

	// 元素个数不确定，可以使用...，由编译器推断
	var data2 = [...]int{1, 2, 3}
	fmt.Println(len(data2))

	// 元素取值的索引规则与java等一致
	fmt.Println(data2[0])

	// 可以使用for循环，foreach循环等遍历数组
	for k, v := range data1 {
		fmt.Printf("第%d个元素为%.2f\n", k, v)
	}
}
