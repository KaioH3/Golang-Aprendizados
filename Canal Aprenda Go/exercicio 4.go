// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type umtipo float32 //cria um tipo, em que o tipo subjacente é o float32

var x umtipo

func main() {
	fmt.Printf("%v do tipo %T\n", x, x)
	x = 42
	fmt.Printf("%v do tipo %T", x, x)
}
