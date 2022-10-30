// You can edit this code!
// Click here and start typing.
package main

import "fmt"

/*
var x int //--> declara a variavel e o tipo da variavel(declaração)

func main() {
	x = 10 //--> inicializa a variavel(inicialização)
	fmt.Println(x)

	x = 20 //--> atribui um valor(atribuição)
	fmt.Println(x)

	y := 15 //--> faz todo o processo de cima(declarou, inicializou e atribuiu), por isso (:=) se chama operador curto de declaração, e sempre que possivel, use-o
		// so use var variavel para package-level scope
	fmt.Println(y)
}
*/

// valor zero
var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
