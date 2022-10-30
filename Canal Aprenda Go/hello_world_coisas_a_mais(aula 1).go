// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//isto e um comentario
	//numerobytes, erros := fmt.Println("Hello, world!", "ola galera", 100)
	//fmt.Println(erros) desta primeira forma o programa nao executa, a linguagem go precisa que o programa utilize todas as variaveis declaradas
	//por que os principios da linguagem go é ser eficiente

	//numerobytes, erros := fmt.Println("Hello, world!", "ola galera", 100)
	//fmt.Println(numerobytes, erros) assim mostra o numero de bytes e o numero de erros, que sao os valores que a funcao println retorna

	//_, erros := fmt.Println("Hello, world!", "ola galera", 100) utilizando o blank identifier, reprensetado pelo underline(underscore) para descartar a variavel não ulizada
	//fmt.Println(erros)
	x := 16
	y := "string"
	z := true
	fmt.Println(x, y, z)

}
