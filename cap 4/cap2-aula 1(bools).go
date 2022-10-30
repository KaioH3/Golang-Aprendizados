// booleanos
package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value
	x = (10 < 0)
	fmt.Println(x)
	d := 0 < 10 // ou (0 < 10)
	y := 20 <= 0
	w := 100 != 10
	z := 100 == 2
	fmt.Println(d)
	fmt.Println(y)
	fmt.Println(w)
	fmt.Println(z)
}
