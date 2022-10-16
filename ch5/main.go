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
	fmt.Println("Nota = ", total/5)
}
func ex3() {
	var x [5]float64
	x[0] = 34
	x[1] = 56
	x[2] = 94
	x[3] = 23
	x[4] = 18

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(x, " = ", total)
	fmt.Println("Nota = ", total/float64(len(x)))
}

func ex4() {
	x := [5]float64{10, 23, 23, 42, 32}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(x, "Total ", total)

}

func ex5() {
	x := [5]float64{
		10,
		23,
		23,
		42,
		32,
	}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(x, "Total ", total)

}

func ex6() { //Slices
	x := make([]float64, 5)
	fmt.Println(x)
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

}

func ex7() { //Slices
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func ex8() { // MAPS
	// var x map[string]int --> x is a map of strings to ints
	x := make(map[string]int)
	x["key"] = 10
	x["key2"] = 20
	fmt.Println(x)
	fmt.Println(x["key"])
	fmt.Println(x["key2"])
	delete(x, "key2")
	fmt.Println(x)
}

func ex9() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Flourine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
}

func ex10() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitorgen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Flourine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

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
	ex9()
	ex10()
}
