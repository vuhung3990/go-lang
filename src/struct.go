package main

import (
	"fmt"
	"math"
)

// Object
type Student struct {
	name  string
	age   int
	class string
}

// method of object
// REMEMBER : (student Student)
func (student Student) who() string {
	return student.name
}

// method single
func who(student Student) string {
	return student.name
}

// INTERFACE
// object cycle
type Cycle struct {
	radian float64
}
// object rectangle
type Rect struct {
	width, height float64
}
// area of 2 object
func (r Rect)area() float64 {
	return r.height * r.width
}
func (c Cycle)area() float64 {
	return c.radian * c.radian * math.Pi
}

type Geometry interface {
	area() float64
}
func measure(g Geometry) float64 {
	return g.area()
}

func main() {
	// sample object
	fmt.Println(Student{"hung", 20, "509cnt"});
	fmt.Println(Student{class: "hung", age: 20, name: "509cnt"});

	s := Student{name: "Brother", age: 23}
	fmt.Println(s.who());
	fmt.Println(who(s));

	// interface
	r := Rect{width:3, height:4}
	c := Cycle{radian:5}
	fmt.Println(measure(r));
	fmt.Println(measure(c));
}
