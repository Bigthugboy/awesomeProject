package main

import (
	"fmt"
)

func main() {
	y := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(y[2:5])

	z := make([]int, 3, 9)
	fmt.Println(z)

	B := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(B)

	fmt.Println(add(50, 3))

	q := 1000
	increment := func() int {
		q++
		return q
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(factorial(5))

	//for i := 0; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println("even")
	//	} else {
	//		println("odd")
	//	}
	//	//fmt.Println(i)
	//}
	// print even and odd number

	//for i := 0; i <= 100; i++ {
	//	if i % 3 == 0 {
	//		fmt.Println(i)
	//
	//	}
	//
	//}
	// print number btw 1 - 100 that is divisible by 3

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else if i%15 == 0 {
			println("fizzbuzz")

		} else {
			println(i)

		}

	}
	var x [5]int
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

}

func add(a, b int) any {
	return a + b
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
