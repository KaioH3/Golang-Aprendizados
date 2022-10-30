// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const ( // iota é uma forma de autoincrmento e serviria meio que para criar um id
	a = 10
	b = 0
	c = iota
	d
	e
	f = 10
	g = iota
	h = 10
)

const ( // fazendo isso, o indice(ou incremento) retorna ao zero
	w = iota
	y
	x
	z
	j
	k
)

const (
	_ = iota // usa-se o underline para que o iota increment descarte esse valor(pule um valor)
	p        //= iota
	ll
	_
	_
	qq
)

const (
	uu = iota     //iota(0) = 0
	tt = iota + 1 //iota(1) = 1+1
	zz = iota * 3 //iota(2) = 2*3 = 6
	ff = iota - 2 //iota(3) = 3-1 = 1
)

// iota pode começar de outros numeros sem ser o zero
const (
	pri = iota + 3 //iota(0) = 0+3
	sec
	terc
	quart
)

// label encoder simples
type Label uint8

const (
	small Label = iota
	medium
	large
	extralarge
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, d, e, f, g, h)
	fmt.Println(w, y, x, z, j, k)
	fmt.Println(ll, qq)
	fmt.Println(ll, qq)
	fmt.Println(uu, tt, zz, ff)
	fmt.Println(pri, sec, terc, quart)
	fmt.Println(small, medium, large, extralarge)
}
