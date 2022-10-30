package main

import (
	"fmt"
)

func main() {
	var x, d, oo, tt int = 10, 2, 3, 4 //variaveis multiplas que só podem ser do tipo int
	var q, w, e, r = 10, 23, 45.1, "aceita qualquer tipo"
	p, i, y, z := 1, 3.4, "aceita qualquer tipo", 5

	var ( // também podem ser declaradas desta forma
		ii int
		ç  int    = 1
		c  string = "hello"
		ll        = 10.1
	)

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[float32]int{1.0: 1, 1.1: 2, 1.2: 3, 1.3: 4}

	priln := fmt.Println
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	//fmt.Println()
	priln(9)
	priln("sei la irado demais")

	fmt.Println("o map de nome b(1.1): ", b[1.1])

	fmt.Println("as variaveis multiplas do tipo int: ", x, d, oo, tt)

	priln("as variaveis multiplas de qualquer tipo(parte 1):\t", q, w, e, r)
	priln("as variaveis multiplas de qualquer tipo(parte 2):\t", p, i, y, z)

	priln("as variaveis multiplas de qualquer tipo(parte 3)(bloco):\t", ll, ii, ç, c)
}
