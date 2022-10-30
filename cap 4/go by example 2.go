// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

const n = 500000000

func main() {
	const d = 3e20 / n
	fmt.Println("d: ", d, "\nn: ", n)

	fmt.Println("int64(d)-->", int64(d))

	fmt.Println(math.Sin(n)) // math.Sin pede um argumento do tipo float64, portanto, constantes não tem um tipo até serem usadas
}
