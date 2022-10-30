// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
)

var x int8

// (um apelido pra uint8)byte = uint8
// (strings são feitas de runes)rune = int32(signed int)
// regra geral: use float64
func main() {
	x = 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//####################### segunda parte
	fmt.Println("\nsegunda parte:")
	a := "e" // usa 1 bytes(8 bits)
	b := "é" // usa 2 bytes(8 bits)
	c := "道" // usa 3 bytes(8 bits)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)

	fmt.Println("\nrunas convertidas para bytes:")
	d := []byte(a) //converte o tipo para um slice de um byte sei lá
	e := []byte(b)
	f := []byte(c)
	// irá retornar slices de bytes(8bits), por isso o tipo que ele retorna é tipo uint8(unsigned int byte)
	fmt.Printf("valor de a: %v e seu tipo:%T\nvalor de b: %v e seu tipo: %T\nvalor de c: %v e seu tipo:%T\n", d, d, e, e, f, f)

	// terceira parte(default)
	fmt.Println("\nterceira parte:")
	i := 10   //default --> int
	p := 10.1 //default --> float64
	fmt.Printf("valor: %v de tipo: %T\n", i, i)
	fmt.Printf("valor: %v de tipo: %T\n", p, p)

	// pode atribuir assim
	i = 110
	p = 110
	fmt.Printf("valor: %v de tipo: %T\n", i, i)
	fmt.Printf("valor: %v de tipo: %T\n", p, p)

	// para ver qual é o OS e a arquitetura do processador
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
