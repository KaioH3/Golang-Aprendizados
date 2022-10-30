// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	slices := make([]string, 5)
	fmt.Println("(primeiro print)", "empty", slices)

	slices[0] = "um(zero)"
	slices[1] = "dois"
	slices[2] = "3"
	fmt.Println("(primeiro print parte 2)", "de empty pra semi preenchido", slices)

	fmt.Println("segundo print-->", slices[1])

	fmt.Println()

	// for i := 0; i < 6; i++ { --> assim ele irá de 0 a 5(passa 6 vezes pelo codigo)
	for i := 0; i < 5; i++ { //--> aqui ele passa pelo codigo 5 vezes
		// fmt.Println("(terceiro print)", slices[i])
		fmt.Printf("(terceiro print) de numero %v %v\n", i, slices[i])
	}

}
