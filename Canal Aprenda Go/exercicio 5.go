// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type umtipo float32 //cria um tipo, em que o tipo subjacente é o float32

var x umtipo

func main() {
	// fmt.Printf("%v do tipo %T\n", x, x)
	// ou
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v do tipo %T\n", x, x)

	// segunda parte
	// transformando o tipo de x para o seu tipo subjacente
	y := float32(x)

	fmt.Printf("convertendo %T para o tipo(%T) com o mesmo valor de x: %v", x, y, y)
}
