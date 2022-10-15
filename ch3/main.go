package main

import "fmt"

func ex1() {
	var x string = "Hello, World"
	fmt.Println(x)
}

func ex2() {
	var x string
	x = "First"
	fmt.Println(x)
	x = "Second"
	fmt.Println(x)

}

func ex3() {
	var x string
	x = "First"
	fmt.Println(x)
	x = x + " Second"
	// x += " Second"
	fmt.Println(x)
}

func ex4() {
	var x string = "Hello"
	var y string = "World"
	fmt.Println(x == y)
	x = "Hello"
	y = "Hello"
	fmt.Println(x == y)

}

func ex5() {
	x := "Hello from x"
	fmt.Println(x)
	var y = "World from y"
	fmt.Println(y)
	z := 5
	fmt.Println(z)
}

func ex6() {
	x := "Max"
	fmt.Println("My dog's name is", x)
	name := "Max"
	fmt.Println("My dog's name is", name)
	dogName := "Max"
	fmt.Println("My dog's name is", dogName)
}

func ex7() {
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)
}

func ex8() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
}
