// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// slice("[]") é como se fosse uma lista do Python
type lista_nomes []string // crio um tipo chamado "lista_nomes" com tipo subjacente: slice do tipo string

func main() {
	var lista_de_nomes lista_nomes
	lista_de_nomes = append(lista_de_nomes, "primeiro nome") // first argument to append must be a slice; have "primeiro nome" (untyped string constant)
	lista_de_nomes = append(lista_de_nomes, "segundo nome")
	lista_de_nomes = append(lista_de_nomes, "terceiro nome")

	for indice, nome1 := range lista_de_nomes {
		indice++
		fmt.Printf("%vº\t%v\n", indice, nome1)
	}
}
