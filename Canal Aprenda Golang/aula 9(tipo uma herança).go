package main

import "fmt"

/*type Categoria struct {
	Nome string
	Pai  *Categoria // para se declarar uma struct que se auto contenha, deve se passar o ponteiro para a struct
}

// metodo que ve a categoria é a pai(parent)("a diferença do metodo para a função é que metodos estão atrelados a structs,")
func (c Categoria) TemParent() { //bool {
	// return c.Pai != nil
	if c.Pai == nil {
		fmt.Println("Não tem pai")
	} else if c.Pai != nil {
		fmt.Println("Tem pai")
	}
}

func main() {
	cat := Categoria{Nome: "Categoria 1"}
	cat.TemParent()

	fmt.Println("categoria com pai?")
	catComPai := Categoria{Nome: "Categoria 1", Pai: &Categoria{Nome: "Pai"}} // faz o referenciamento(apontamento) da memoria e cria a struct categoria com nome Pai
	catComPai.TemParent()
	if !cat.TemParent() {
		fmt.Printf("Não tem pai!")
	}

}*/

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
} // agora, quando eu for

func main() {
	pe := Pessoa{"Tiago", "Temporin", 34, true, "0.00.0.0.0.0"}
	pe2 := Pessoa{"Jorgin", "Temporao", 204, true, "3.01.0.0.0.0"}
	// ao inves de eu executar tudo isso fmt.Println("Nome:", p.Nome, "\nSobrenome:", p.Sobrenome),
	// fmt.Println(pe) // eu só mostro o struct que ele executa o metodo String()
	fmt.Println(pe, "e tenho o cpf:", pe.cpf) // ainda posso usar as variaveis do structs

	fmt.Println("pessoa 2:", pe2, "e tenho o cpf:", pe2.cpf)
}
