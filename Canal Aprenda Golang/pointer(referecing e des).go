package main

import (
	"fmt"
)

func main() {
	// refereciando
	// basicamente o pointer aponta para um local na memória(se referencia a local na memoria)
	// basicamente armazena o endereço do local da memoria

	a := 45

	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a) // mostra o endereço da memoria

	// sobre o processo de referenciamento(ou apontamento sei la)
	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

	var b = &a // apontando ou referenciando a local da memoria em que a variavel "a" esta armazenada

	fmt.Println("b: ", b)

	// "desrefereciando"

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case

	// var b = &a o local da memoria tava sendo referenciado aqui, certo?
	// então agora, o processo para mostra o que tem nesse endereço, é chamado de desreferenciamento
	fmt.Println(b)  // 0x20818a220 --> o endereço do local da memoria
	fmt.Println(*b) // 43 --> (dereferencing)(desreferenciando)"decifrando" o que está no local da memoria
}
