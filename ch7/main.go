package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func ex1() {
	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 5}

	fmt.Println(r.area())
	fmt.Println(c.area())
}

type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is, " + p.Name)
}

func (p *Person) setName(name string) {
	p.Name = name
}

type Android struct {
	Person
	Model string
}

func ex2() {
	a := new(Android)
	a.setName("Gabu Android")
	a.talk()

	p := new(Person)
	p.setName("Gabu Pessoa")
	p.talk()
}

// func totalArea(circles ...Circle) float64 {
// 	var total float64
// 	for _, c := range circles {
// 		total += c.area()
// 	}
// 	return total
// }

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

type Multishape struct {
	shapes []Shape
}

func (m *Multishape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func ex3() {
	c := Circle{0, 0, 5}
	c1 := Circle{0, 0, 2}
	r := Rectangle{0, 0, 10, 10}
	r1 := Rectangle{0, 0, 10, 10}

	fmt.Println(totalArea(&c, &c1, &r, &r1))

	multishape := Multishape{
		shapes: []Shape{
			&c,
			&r,
		},
	}

	fmt.Println(multishape.area())

}

func main() {
	ex1()
	ex2()
	ex3()
}
