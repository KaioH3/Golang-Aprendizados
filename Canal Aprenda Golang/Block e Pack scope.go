package main

import "fmt"

var x = 0

func incrementBlockScope() int {
	y := 0
	y++
	return y
}

func incrementPackScope() int {
	x++
	return x
}

func incretsBS() {
	fmt.Println(incrementBlockScope())
	fmt.Println(incrementBlockScope())
	fmt.Println(incrementBlockScope())
	fmt.Println(incrementBlockScope())
	fmt.Println(incrementBlockScope())
	fmt.Println(incrementBlockScope())

}

func incretsPS() {
	fmt.Println(incrementPackScope())
	fmt.Println(incrementPackScope())
	fmt.Println(incrementPackScope())
	fmt.Println(incrementPackScope())
	fmt.Println(incrementPackScope())
	fmt.Println(incrementPackScope())
}

func wrapper() func() int {
	z := 0
	return func() int {
		z++
		return z
	}
}

func main() {
	w := 0
	incrementAnonymous := func() int {
		w++
		return w
	}

	fmt.Println("incrementAnonymous", incrementAnonymous())

	incrementWrapper := wrapper()
	fmt.Println("incrementWrapper", incrementWrapper())

	fmt.Println("\nBlock scope: ")
	incretsBS()
	fmt.Println("\nPackage scope: ")
	incretsPS()

}
