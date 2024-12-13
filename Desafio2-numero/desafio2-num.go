/*
"Programa para contar de 1 a 100, se for multiplo de 3, aparece pin, se for multiplo de 5, mostra pan,
se não aparece o numero normal.
*/
package main

import "fmt"

func main() {
	for i:= 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Println("pin")
		} else {
			if i%5 == 0 {
				fmt.Println("pan")
			} else{
				fmt.Println(i)
			}
		} 
	}
}