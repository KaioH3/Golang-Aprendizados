// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// strings são imutaveis, entao para alterar essa string teria que criar outra
// uma string é uma slice de bytes
func main() {
	s := "H0é¬ä道"
	// sb := []byte(s) //slice de bytes
	// fmt.Printf("%V", sb, sb)
	/*
		for _, v := range sb {
			fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v) // valor do byte - tipo - unical endpoint - hexadecimal
		}*/

	// agora usando o range da string(vai retornar caractere por caractere)
	for _, v := range s {
		fmt.Printf("%v - %b - %T - %#U - %#x\n", v, v, v, v, v) // valor do byte - tipo - unical endpoint - hexadecimal
	}

	// retorna byte por byte
	// então aqui vai retornar caracteres e as informações destes, de 8 em 8 bits
	// ou melhor, irá retornar caracteres sem ultrapassar 1 byte(8bits)
	fmt.Println("\nretorna byte por byte")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}
}
