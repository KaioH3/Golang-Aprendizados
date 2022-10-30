// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// cria o tipo "hotdog"
type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", x)

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", x)

	// b = x --> nao pode fazer isso: cannot use x (variable of type int) as type hotdog in assignment
}
