// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	slice1 := []string{"primeiro", "segundo", "terceiro", "quarto", "quinto", "sexto"}
	// range para slice^ retorna (indice, elemento)

	string1 := "Hoje, tem! 知道"
	map1 := map[int]string{1: "um", 2: "dois", 3: "tres", 4: "quatro", 5: "cinco", 6: "seis"}

	for indice, elemento := range slice1 {
		fmt.Printf("slice[%v] = %v\n", indice, elemento)
	}
	seila := "#####################################################################\n"
	fmt.Printf(seila)
	for indice, elemento := range string1 {
		fmt.Printf("string[%v] = %v\truna:(%v)\n", indice, string(elemento), elemento) // percorre a string, mostra a stransfora
	}
	fmt.Printf(seila)
	for chave := range map1 {
		fmt.Printf("map1[%v]\n", chave) // percorre a string, mostra a stransfora
	}
}
