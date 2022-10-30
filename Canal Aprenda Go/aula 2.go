// You can edit this code!
// Click here and start typing.
// go doc fmt Printf --> comando no prompt para ver o que significa o comando Printf do modulo "fmt"
package main

import (
	"fmt"
)

// o gopher(:=) só funcionar dentro de block codes, como o debaixo
//var z := "ola bom dia" --> não funciona aqui
//var z = "ola bom dia" --> funciona aqui

func main() {
	//operador de declaração
	x := 5 + 3
	y := "operador de declaração"
	fmt.Println(x, y)
	//fmt.Println(x, y, z)

	//operador de atribuição
	x = 10 + 10
	y = "operador de atribuição"
	fmt.Println(x, y)

	//valor booleano
	w := 10 == 10
	// y = 10 == 9 --> nao posso fazer isso, ja que o y foi declarado inicialmente como string
	z := 10 == 9
	fmt.Println(w, 'z', z)
}
