package main

import "fmt"

func main() {
	nomes := []string{"Seilha", "Seila", "Martha", "Junior", "Thiago"} // slices são array dinamincos que não tem ou não precisam ter um tamanho definido
	/*for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	} pode ser feito assim*/
	/*ou assim:
	for _, nome := range nomes {
		fmt.Println("sem o indice", nome)
	}*/
	// agora como fazer um "while"
	var i int
	for i < len(nomes) { // semelhante ao laço while(enquanto i for menor que o tamanho de "nomes", faça isso)
		// e como "i" começa em zero, colocar <= não funciona
		fmt.Println(nomes[i])
		i++
	}

	/*loop infinito
	for {faça isso}*/
}
