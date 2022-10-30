package main

import "fmt"

/*
func main() {
	nomes := []string{"Seilha", "Seila", "Martha", "Junior", "Thiago"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Neurton") // função built-in que adiciona elementos no fim do slice, e retorna um novo slice com esse elemento
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Neto")
	fmt.Println(nomes, len(nomes), cap(nomes))

	//completando toda a capacidade do slice
	nomes = append(nomes, "Rafael")
	nomes = append(nomes, "Susanda")
	nomes = append(nomes, "Supara")
	nomes = append(nomes, "Suparanda")
	fmt.Println(nomes, len(nomes), cap(nomes))
}*/

// função make
/*func main() {
	// a função make aloca e inicializa valores do tipo slice, map ou channel
	nomes := make([]int, 3, 5)                 //cria um slice com 3 valores-zero e capacidade 5(que por baixo dos panos é um array com capacidade 10)
	fmt.Println(nomes, len(nomes), cap(nomes)) // [0 0 0] 3 5

}*/

/*func main() {
	// map é tipo um dicionario do python(e é tipo hashmap do java)
	idades := make(map[string]uint8) // aloca na memoria e inicializa um map com chaves em string e valores em uint8(uint8 porque idade não passa de 127 anos na vida real...)
	//ou assim:
	var names = map[string]uint8 {
		"John": 49,
		"Jane": 36,   // a ultima virgula é obrigatória
	}
	idades["Tiago"] = 31
	idades["Daniela"] = 36
	idades["Maria"] = 23

	fmt.Println(idades) //  map[Daniela:36 Maria:23 Tiago:31], maps não retornam seus valores de forma ordenada
}*/

//tipo struct
// primeira letra maiscula nas variaveis que voce queira que estejam disponiveis para outros packages(pacotes)
// por exemplo, se tiver uma struct que tem um cpf, que precisa de mais segurança, e você cria um programa que pega registros de usuarios,
// através de uma função Scan(), você coloca os comandos de leitura do stind no package da função Scan, para que o usuario final nçao possa acessar os cpfs tao facilmente
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

type Categoria struct {
	Nome string
	Pai  *Categoria // para se declarar uma struct que se auto contenha, deve se passar o ponteiro para a struct
}

func main() {
	/* pode ser feito assim:
	var p Pessoa
	p.Nome = "Tiago"
	p.Sobrenome = "Temporin"*/

	// ou assim:
	/*p := Pessoa{
		Nome:      "Tiago",
		Sobrenome: "Temporin",
		Idade:     34,
		Status:    true,
	}*/
	// ou ainda, assim:
	//p := Pessoa{"Tiago", "Temporin", 34, true} // atribui valores de acordo com a ordem em que a struct foi cria
	p := Pessoa{"Tiago", "Temporin", 34, true}
	// porém desta forma não é tão bom, já que se colocarem um valor a mais na struct no meio dela, vai bagunçar tudo

	// fmt.Println(p) //imprime tudo
	fmt.Println("Nome:", p.Nome, "\nSobrenome:", p.Sobrenome)
}
