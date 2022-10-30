package main

import "fmt"

func verPar(x int) {
	fmt.Println(x)
	x = x % 2

	if x == 1 {
		fmt.Println("impar")
	} else {
		fmt.Println("par")
	}
	fmt.Println("######\n")
}

func main() {
	//x := 13 % 3

	for i := 1; i < 70; i++ {
		verPar(i)
		/*if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}*/
	}
}
