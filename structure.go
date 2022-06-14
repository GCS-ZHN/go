package main

import "fmt"

/*
在C语言的结构体中，其是内存上连续的一段数据，但
和数组不同的事，可以使用不同的数据类型。go和C类似，
并不直接支持OOP的高级特性，但可以使用结构体来实现
数据的结构化和OOP。其定义方式与C也相似，使用type和
struct关键字来声明。
*/
type Book struct {
	title   string
	author  string
	subject string
	id      int
	price   float64
}

/*
如前所述，可以通过函数指定接受者来实现方法
这与高级语言的方法实现this指针异曲同工
*/
func (this Book) getDescription() {
	fmt.Printf("《%s》由%s创作，所属主题是%s，售价%.2f，编号%d\n", this.title, this.author, this.subject, this.price, this.id)
}

/*
结构体成员可以通过.来访问，这与C语言一样
不过，go里面及时是结构体指针，也是用.来访问
即.同时具有C中->的解引用作用
*/
func (this Book) new(title, author, subject string, price float64, id int) Book {
	this.title = title
	this.author = author
	this.subject = subject
	this.price = price
	this.id = id
	return this
}

func main() {
	/* 创建结构体实例并初始化
	可以按块初始化，也可以逐个成员初始化
	*/
	var book0 = Book{"python程序设计", "张洪宁", "编程", 0, 48}
	fmt.Printf("%x\n", book0)
	book0.getDescription()

	var book Book
	// 结构体本身其实就是一个指针，初始化前为空
	fmt.Printf("%x\n", book)

	/* 调用方法时，隐式调用 this = book进行按值传递
	在new后，返回的this != book，因此需要赋值
	*/
	book = book.new(
		"深入浅出Pytorch--从模型到源码",
		"张校捷",
		"编程；AI",
		89,
		1)

	// 初始化后，指针地址发生变化了
	fmt.Printf("%x\n", book)
	book.getDescription()

	// 结构体指针访问成员的方式不变
	// 因此暂未发现结构体指针的意义
	fmt.Printf("%x\n", &book)
	(&book).getDescription()
}
