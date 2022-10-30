// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// constantes só são definidadas quando forem usadas, diferente de var x = 10

// const x = 10
const (
	y = 15
	d = 10.1
)

// pode declarar consts assim

var (
	o = 19
	p = 20.1
) // pode declarar variaveis assim

func main() {
	// o = y
	p = y
	fmt.Printf("%T", p)

}
