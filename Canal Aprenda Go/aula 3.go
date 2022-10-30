// You can edit this code!
// Click here and start typing.
package main

import "fmt"

/* nao funciona
func main() {
	y := 10 --> o erro, a variavel y esta declara dentro deste code block
	qualquer(y)
}

func qualquer(x int) {
	fmt.Println(y) --> o erro...
	fmt.Println(x)
}
######################################
*/

// desta forma:
var y = 10 //--> desta forma, a variavel se torna package-level scope e não se pode usar o gopher(:=)
func main() {
	y := 10
	qualquer(y)
}

func qualquer(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
