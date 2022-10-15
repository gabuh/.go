package main

import "fmt"

func ex1() {
	i := 0
	for i <= 10 {
		i++
		fmt.Println(i)
	}
}

func ex2() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
func ex3() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}

	}
}

func ex4() {
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		default:
			fmt.Println("none")
		}
	}
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()

}
