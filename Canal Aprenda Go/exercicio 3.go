// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var x = 42           //int
var y = "James bond" //string
var z = true         //bool

func main() {
	fmt.Printf("%T\n%T\n%T\n", x, y, z)
	// s := fmt.Sprintln(x, y, z)
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Print(s)
}
