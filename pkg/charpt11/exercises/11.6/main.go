// float64 is necessary as input to math.Sqrt()
package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type Point3 struct {
	X, Y, Z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Polar struct {
	R, T float64
}

func (p Polar) Abs() float64 { return p.R }

func main() {
	var m Magnitude
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	// 这里m是一个接口变量，它是一个抽象。 把具体的变量的值赋给它，它就能实现多态了。
	// 所谓多态，就是不论具体是什么，我都对m调用Abs，如果m具体对应的是某个类型，就执行某个类型具体的Abs
	m = p1
	fmt.Printf("The length of the vector p1 is: %f\n", m.Abs())

	p2 := &Point{4, 5}
	m = p2
	fmt.Printf("The length of the vector p2 is: %f\n", m.Abs())

	p1.Scale(5)
	m = p1
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", m.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f", p1.X, p1.Y)
}

/* Output:
The length of the vector p1 is: 5.000000
The length of the vector p2 is: 6.403124
The length of the vector p1 after scaling is: 25.000000
Point p1 after scaling has the following coordinates: X 15.000000 - Y 20.000000
*/
