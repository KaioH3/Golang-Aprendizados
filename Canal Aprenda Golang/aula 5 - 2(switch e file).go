// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"os"
)

func le_arquivo() {
	file, err := os.Open("hello.txt")
	fmt.Println("file ", file)
	if err != nil {
		fmt.Println("deu cao")
		log.Panic(err)
	}

	data := make([]byte, 100) // cria um array de bytes com 100 bytes no total, porque a função que chamamos cria um array

	if _, erroDoRead := file.Read(data); erroDoRead != nil { // primeiro retorna a quantidade de bytes que ele leu, e o segundo retorno é se deu erro
		log.Panic(erroDoRead)
	}
	// fmt.Println(erroDoRead) --> nao funciona justamente porque a variavel "erroDoRead" é declarada só dentro do escopo do if
	fmt.Println(string(data)) // o file.Read retorna um array de bytes, portnto, é necessario transforma-los em string
}

var (
	cara, coroa int
)

// funcionamento do switch case
func lancaMoeda(lado string) {
	switch lado { // a variavel "lado" do tipo string
	case "cara": // caso o argumento("lado") seja "cara",
		cara++ // incremente a variavel(package level) cara

	case "coroa": // se for coroa, faz isso:
		coroa++

	default: // se não for nem cara, nem coroa, usa o default(else sem condição)
		fmt.Println("caiu em pé")
	}

}

func main() {
	a, b := 15, 15
	/*if a > b {
		fmt.Println("a é maior que b")
	} else if a < b {
		fmt.Println("a é menor que b")
	} else {
		fmt.Println("a e b são iguais")
	}
	le_arquivo()*/
	switch { // esse switch é a mesma coisa do amontoado de if e elses de cima
	case a > b:
		fmt.Println("a é maior que b")
	case a < b:
		fmt.Println("a é menor que b")
	default:
		fmt.Println("a e b são iguais")
	}
}
