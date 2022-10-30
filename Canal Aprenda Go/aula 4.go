// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// formas de declaração de variaveis
// var x = 10
// ou var x string = "10"
//Type 	        Size 	        Range
//float32 	32 bits 	-3.4e+38 to 3.4e+38.
//float64 	64 bits 	-1.7e+308 to +1.7e+308.
// assim: var x float32 = 10
//ou assim: var x float32 = 10.0001

// ou var x = "string"
// ou x := "string"
// var x float32 = 10.0

var x int //--> ou ate assim, porém, desta forma, a atribuicao só pode ser feita dentro de um code block

func main() {
	// x = 10.2 --retorna um erro, se a variavel for tipo int, ja que a golang é fortemente tipada
	// caso nada seja atribuido a variavel tipo int, ela retornara o numero zero(0)
	x = 10                     //--> declarando dentro do block code
	fmt.Printf("%v, %T", x, x) //para imprimir a variavel e o tipo dela
}
