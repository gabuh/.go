package main

import "fmt"

func ex1() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func ex2() {
	var x [5]float64
	x[0] = 23
	x[1] = 23
	x[2] = 23
	x[3] = 23
	x[4] = 23

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(x, " = ", total)
}

func main() {
	ex1()
	ex2()
}
