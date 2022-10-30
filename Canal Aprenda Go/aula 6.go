// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// strings:
	// x := "ola, \tbom dia\n\"tudo bem\"" //--> isso é uma string interpretada(interpreted string literals),
	// ou seja, o compilador irá ler estes comando --> \t: tab; \n: newline; \: carctere de escape

	// agora, com o acento grave(``) sendo usados como se fossem as aspas duplas(""), o compilador vai ler como raw string(raw string literals)
	x := `"ola, \tbom // -->
	dia\n
	
	\"tu
	do be
	m\
	"
	"`
	fmt.Println(x)
}
