// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var i uint16
	i = 65535 // tamanho limite do tipo uint16 - 1(65536)
	fmt.Println(i)
	i++            // adicionando um, overflow, e a variavel vai a zero
	fmt.Println(i) // i = 0
	i++
	fmt.Println(i) // i = 1
}
