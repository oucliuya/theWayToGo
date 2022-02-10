package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (s *Simple) Set(v int) {
	s.value = v
}

type RSimple struct {
	i int
	j int
}

func (r *RSimple) Get() int {
	return r.j
}

func (r *RSimple) Set(v int) {
	r.j = v
}

// 这个练习核心是fI函数，它接受一个Simpler接口类型，那实际入参可能是不同的实现了Simpler的类型。fI内部通过 ```.(type)```方法来获取实际入参的类型，并根据具体类型执行具体操作
func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s))
	var r RSimple
	fmt.Println(fI(&r))
}
