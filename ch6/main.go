package main

import "fmt"

func ex1() {
	xs := []float64{89, 79, 56, 87, 45}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}

// func everage(xs []float64) float64{
// 	panic("Not Implemented")
// }

func everage(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func ex2() {
	xs := []float64{89, 79, 56, 87, 45}
	fmt.Println(everage(xs))
}

// Stack call of functions main -> f1 -> f2 -> f1 -> main
// func main(){
// 	fmt.Println(f1())
// }
// func f1() int{
// 	return f2()
// }
// func f2() int{
// 	return 1
// }

func ex3_f2() (r int) {
	r = 1
	return
}
func ex3_f() (int, int) {
	return 5, 4
}

func ex3() {
	x, y := ex3_f()
	fmt.Println(x, " ", y)
	// x, err := f()
}

//func Println( a ...interface{} ) (n int, err error )

func ex4_f(args ...int) (total int) {
	total = 0
	for _, v := range args {
		total += v
	}
	return
}

func ex4() {
	fmt.Print("This variadic f must be 6 --> ")
	fmt.Println(ex4_f(1, 2, 3))
	//passing Slices
	xs := []int{1, 3, 5}
	fmt.Println("This variadic f + slice must be 9 --> ", ex4_f(xs...))

}

func ex5_f() {
	add := func(args ...int) (total int) {
		for _, v := range args {
			total += v
		}
		return
	}
	fmt.Println("Closure Variadic Function Must be 56 --> ", add(55, 1))
}
func ex5() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("Closure Must be 2 --> ", add(1, 1))
	ex5_f()
}

func ex6() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("Closure increment --> ", increment())
	fmt.Println("Closure increment --> ", increment())
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()

}
