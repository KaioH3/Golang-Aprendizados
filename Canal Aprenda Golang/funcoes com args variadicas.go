// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// slice("[]") é como se fosse uma lista do Python
type lista_nomes []string // crio um tipo chamado "lista_nomes" com tipo subjacente: slice do tipo string

func imprime_slice() {
	var lista_de_nomes lista_nomes
	lista_de_nomes = append(lista_de_nomes, "primeiro nome") // first argument to append must be a slice; have "primeiro nome" (untyped string constant)
	lista_de_nomes = append(lista_de_nomes, "segundo nome")
	lista_de_nomes = append(lista_de_nomes, "terceiro nome")

	for indice, nome1 := range lista_de_nomes {
		indice++
		fmt.Printf("%vº\t%v\n", indice, nome1)
	}
}

type maratona_pessoa struct {
	nome        string
	idade       int
	quilometros int
	participou  bool
}

func imprime_struct(competidor maratona_pessoa) { // função que tem como argumento, o tipo struct que tem como tipos subjacentes, o conjunto de string, int e bool
	if competidor.participou {
		fmt.Println(competidor.nome, "andou", competidor.quilometros, "KM", "com", competidor.idade, "anos de idade")
	} else {
		fmt.Println("O competidor", competidor.nome, "infelizmente não participou")
	}
	fmt.Println()
}

func main() {
	roberto := maratona_pessoa{nome: "Robertao", idade: 29, quilometros: 20, participou: true}
	roberto2 := maratona_pessoa{nome: "Robertao2", idade: 29, participou: true} // tirei "quilometros: 20" então o compilador retorna o zero value(que nesse caso é 0)
	roberto3 := maratona_pessoa{idade: 29, participou: true}                    // tirei "quilometros: 20" e "nome: "Robertao2"
	roberto4 := maratona_pessoa{participou: true}                               // tirei "quilometros: 20" e "nome: "Robertao2" e "idade: 29"
	roberto5 := maratona_pessoa{}                                               // sem nada

	fabio := maratona_pessoa{nome: "Fabin", idade: 55, participou: true, quilometros: 13}   // nao importa a ordem em que se atribui os valores(ordem de "quilometros" foi trocada)
	fabio2 := maratona_pessoa{idade: 55, participou: true, quilometros: 13, nome: "Fabin2"} // repito, nao importa a ordem

	fulano := maratona_pessoa{nome: "Fulano", idade: 43, quilometros: 0, participou: false}

	//imprime_slice()

	imprime_struct(roberto)
	imprime_struct(roberto2)
	imprime_struct(roberto3)
	imprime_struct(roberto4)
	imprime_struct(roberto5)

	imprime_struct(fabio)
	imprime_struct(fabio2)
	imprime_struct(maratona_pessoa{idade: 55, participou: true, quilometros: 13, nome: "Fabin sem variavel atribuida"})

	imprime_struct(fulano)
}

###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// slice("[]") é como se fosse uma lista do Python
type lista_nomes []string // crio um tipo chamado "lista_nomes" com tipo subjacente: slice do tipo string

func imprime_slice() {
	var lista_de_nomes lista_nomes
	lista_de_nomes = append(lista_de_nomes, "primeiro nome") // first argument to append must be a slice; have "primeiro nome" (untyped string constant)
	lista_de_nomes = append(lista_de_nomes, "segundo nome")
	lista_de_nomes = append(lista_de_nomes, "terceiro nome")

	for indice, nome1 := range lista_de_nomes {
		indice++
		fmt.Printf("%vº\t%v\n", indice, nome1)
	}
}

type maratona_pessoa struct {
	nome        string
	idade       int
	quilometros int
	participou  bool
}

func imprime_struct(competidor maratona_pessoa) { // função que tem como argumento, o tipo struct que tem como tipos subjacentes, o conjunto de string, int e bool
	if competidor.participou {
		fmt.Println(competidor.nome, "andou", competidor.quilometros, "KM", "com", competidor.idade, "anos de idade")
	} else {
		fmt.Println("O competidor", competidor.nome, "infelizmente não participou")
	}
	fmt.Println()
}

func main() {
	roberto := maratona_pessoa{nome: "Robertao", idade: 29, quilometros: 20, participou: true}
	roberto2 := maratona_pessoa{nome: "Robertao2", idade: 29, participou: true} // tirei "quilometros: 20" então o compilador retorna o zero value(que nesse caso é 0)
	roberto3 := maratona_pessoa{idade: 29, participou: true}                    // tirei "quilometros: 20" e "nome: "Robertao2"
	roberto4 := maratona_pessoa{participou: true}                               // tirei "quilometros: 20" e "nome: "Robertao2" e "idade: 29"
	roberto5 := maratona_pessoa{}                                               // sem nada

	fabio := maratona_pessoa{nome: "Fabin", idade: 55, participou: true, quilometros: 13}   // nao importa a ordem em que se atribui os valores(ordem de "quilometros" foi trocada)
	fabio2 := maratona_pessoa{idade: 55, participou: true, quilometros: 13, nome: "Fabin2"} // repito, nao importa a ordem

	fulano := maratona_pessoa{nome: "Fulano", idade: 43, quilometros: 0, participou: false}

	//imprime_slice()

	imprime_struct(roberto)
	imprime_struct(roberto2)
	imprime_struct(roberto3)
	imprime_struct(roberto4)
	imprime_struct(roberto5)

	imprime_struct(fabio)
	imprime_struct(fabio2)
	imprime_struct(maratona_pessoa{idade: 55, participou: true, quilometros: 13, nome: "Fabin sem variavel atribuida"})

	imprime_struct(fulano)
}

###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// funções com argumentos variadicos, mais ou menos como tem no Python

func foo(params ...int) { // params is a slice of ints
	fmt.Println(len(params))
}

// foo com laço foor
func foor(params ...int) {
	fmt.Println("tamanho:", len(params))

	for idx, item := range params {
		idx++
		fmt.Printf("%vº item: %v\n", idx, item)
	}
}

// struct recebe valores, mas não necessariamente precisa ter todos os valores atribuidos para ser usado, porém todos os valores são declarados(ficam preenchidos com o valor zero)
type Params struct {
	a, b     int
	operacao string
	inicia   bool
}

func calculadora(p Params) { // recebe um argumento do tipo struct(p Params)
	if p.inicia {
		if p.operacao == "+" {
			Resultado := p.a + p.b
			fmt.Println(p.a, "-", p.b, "=", Resultado)
		} else if p.operacao == "-" {
			Resultado := p.a + p.b
			fmt.Println(p.a, "-", p.b, "=", Resultado)
		}
	}
}

// you can call it without specifying all parameters
//doIt(Params{a: 1, c: 9})

func main() {
	/*foo()
	foo(1)
	foo(1, 2, 3)

	fmt.Println("\nFoo com laço for:")
	foor(1, 2, 3)*/
	calculadora(a:1, b:2)
}

###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################
###############################$#$$$$$###################################$$$$$$######################################

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Params struct {
	a, b     int
	operacao string
	inicia   bool
}

func calculadora(p Params) { // recebe um argumento do tipo struct(p Params)
	if p.inicia {
		if p.operacao == "+" {
			Resultado := p.a + p.b
			fmt.Println(p.a, "-", p.b, "=", Resultado)
		} else if p.operacao == "-" {
			Resultado := p.a + p.b
			fmt.Println(p.a, "-", p.b, "=", Resultado)
		}
	} else {
		fmt.Println("Chamou a função mas não a inicializou")
	}
}

func main() {
	calculadora(Params{}) // Chamou a função mas não a inicializou
	
}
