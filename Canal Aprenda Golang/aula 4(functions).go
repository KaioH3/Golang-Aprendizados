// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

// funções com a primeira letra minuscula só podem ser chamadas neste pacote(package)
// para que possam ser chamadas em outros pacotes, tem que se colocar a primeira letra como maiuscula(export a função)(func FuncaoExportada)
// e ocorre a mesma coisa para variaveis
func hello(nome string) {
	fmt.Println("Hello função", nome, "!")
}

func somar(a int, b int) int { // pega dois inteiros e retorna um inteiro
	return a + b
}

func convertsrtAndSum(a int, b string) int { //(variaveis de entrada) tipo da variavel de retorno
	i, _ := strconv.Atoi(b)
	return a + i
}

func convertsrtAndSumComErro(a int, b string) (total int, err error) {

	i, err := strconv.Atoi(b) // a variavel "i" é a unica que esta sendo criada no momento

	if err != nil {
		fmt.Print("deu cao:\t")
		return
	}
	// total := a + i isso não funciona, pois as variaveis "total", "a" e "i" já foram declaradas a declarar a função
	// então somente atribuimos valores a elas
	total = a + i
	return
}

func ConvertsrtAndSumComErro(a int, b string) (total int, err error) { // função exportada(letra maiuscula)

	i, err := strconv.Atoi(b) // a variavel "i" é a unica que esta sendo criada no momento

	if err != nil { // se erro diferente de nil(null, sem erro), execute os comandos
		fmt.Print("deu cao:\t")
		return
	}
	// total := a + i isso não funciona, pois as variaveis "total", "a" e "i" já foram declaradas a declarar a função
	// então somente atribuimos valores a elas
	total = a + i
	return
}

func main() {

	hello("sei la")
	// somar(1, "2")
	fmt.Println(somar(20, 2))
	fmt.Println(convertsrtAndSum(1, "200"))
	fmt.Println(convertsrtAndSumComErro(1, "200"))
	fmt.Println(convertsrtAndSumComErro(1, "200l")) //--> com erro
	fmt.Println(ConvertsrtAndSumComErro(1, "20"))
}
