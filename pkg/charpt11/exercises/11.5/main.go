package main

import "fmt"

/*
这个例子展示了go语言中一个类型可以实现多个接口。 struct的成员是数据，interface的成员是方法，这样就实现了数据和方法的解耦
如果是java， 我们如果要给两个不同的类A、B定义同一个方法func1，我们要定义一个基类，然后让A、B都继承这个基类，然后在A、B里面
具体定义它们的func1

这里我们定义了两个接口，分别在这两个接口里面定义了Area和Perimeter方法。其实也可以在一个接口里面定义这两个方法。
可见接口是非常灵活可拆分的
*/

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

func (tri *Triangle) Area() float32 {
	return 0.5 * tri.base * tri.height
}

func main() {
	sq := Square{5}
	fmt.Printf("square area is:%f, square primeter is:%f\n", sq.Perimeter(), sq.Area())
	tri := Triangle{base: 3, height: 4}
	fmt.Printf("triangle area is:%f\n", tri.Area())
}
