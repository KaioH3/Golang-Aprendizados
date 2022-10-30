// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var arrayzin1 = [10]int{1, 2}                    // com o comprimento
var arrayzin2 = [...]float64{1.2, 1.3, 1.4, 2.6} // o compilador vai inferir o comprimento(advinhar)

//var arrayzin1 =

// priln := fmt.Println --> da erro aqui(package level)(goopher--> :=): non-declaration statement outside function body
// var priln = fmt.Println --> isso funciona aqui.

func main() {
	arrayzin3 := [5]string{"seila", "seila2", "seila3"} // com total de 5 valores e só 3 valores incializados
	arrayzin4 := [...]bool{true, false, true, false}    // com um total de acordo com a quantidade de valores declarados inicialmente
	arrayzin5 := [6]bool{true, false, true, false}      // com total de 6 valores e só 4 valores inicializados

	// legenda: inf para inferido e decl para declarado
	priln := fmt.Println

	priln("tamanho(lenght)(decl):", len(arrayzin1), "| capacidade", cap(arrayzin1), "| valores do array 1(parte 2):\t", arrayzin1)
	//"tamanho: \t", len(arrayzin), "valores do array:\t",
	// arrayzin1[2] = 3 o tamanho do array é invariavel, entao se ele tem 2 espaços, ele não passa disso

	arrayzin1[1] = 3
	priln("tamanho(lenght)(decl):", len(arrayzin1), "| capacidade", cap(arrayzin1), "| valores do array 1(parte 2):\t", arrayzin1)
	// priln("tamanho: ", "lenght:", len(arrayzin1), "| capacidade", cap(arrayzin1), "| valores do array :\t"

	priln("tamanho(lenght)(inf):", len(arrayzin2), "| capacidade", cap(arrayzin2), "| valores do array 2:\t", arrayzin2)
	priln("tamanho(lenght)(decl):", len(arrayzin3), "| capacidade", cap(arrayzin3), "| valores do array 3:\t", arrayzin3)
	priln("tamanho(lenght)(inf):", len(arrayzin4), "| capacidade", cap(arrayzin4), "| valores do array 4:\t", arrayzin4)
	priln("tamanho(lenght)(decl):", len(arrayzin5), "| capacidade", cap(arrayzin5), "| valores do array 5:\t", arrayzin5)

	// o compilador vai dizer que o array tem esse tamanho e capacidade, porém nao irá inicializar os valores, e vai preencher-los com zero values
	// e.g. false para bools, 0 para ints, e nada(e assim: "" para de acordo com o go style) para strings

	// seila(arrayzin5)
}

// func seila(a int, b int) int {} assim
// ou assim:
// func seila(a []interface{}) {return (fmt.Println("tamanho: ", "lenght:", len(a), "| capacidade", cap(a), "| valores do array :\t", a))}
// nao foi, entao tenta denovo mais tarde
