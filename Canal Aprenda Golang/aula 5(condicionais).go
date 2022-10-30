// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// funcionamento do if sem else:  if(qualquer um: true ou false)-->proxima linha
// funcionamento do if else:  if (false)-->else / if (true)-->pula todos os else's-->proxima linha
// var b, c string = "stored in b", "stored in c" // package scope
// var B, C string = "stored in b", "stored in c" // exported variables

func numeroParSemStatementDireto(num int) {
	// num := 11 % 2
	modulo := 11 % 2
	if modulo == 0 { //ve se o numero é par
		fmt.Println("O número", num, "é par")
		return
	} else {
		fmt.Println("O número", num, "é impar")
	}
}

func numeroParComStatementDireto(num int) {
	// if num:= 0; num % 2 == 0{} é assim declarando e atribuindo diretamente
	if modulo := num % 2; modulo == 0 { //ve se o numero é par
		fmt.Println("O número", num, "é par")
		return
	} else {
		fmt.Println("O número", num, "é impar")
	}
}

func main() {
	/*j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks */
	a, b := 15, 15
	if a > b {
		fmt.Printf("a é maior que b\n")
	} else if a < b {
		fmt.Printf("a é menor que b\n")
	} else if a == b {
		fmt.Printf("a e b são iguais\n")
	}

	numeroParSemStatementDireto(2)
	numeroParSemStatementDireto(3)
	numeroParSemStatementDireto(3031)
	// numeroParSemStatementDireto(03031) dá um numero estranho, deve retornar algum binario nao sei
	// numeroParSemStatementDireto(0x3031) dá um numero estranho também
	fmt.Println()
	numeroParComStatementDireto(2)
	numeroParComStatementDireto(3)
	numeroParComStatementDireto(3031)

}
