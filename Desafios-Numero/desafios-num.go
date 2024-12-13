/*"Programa para mostrar todos multiplos de 3 de 1 a 100."*/
package main

import "fmt"

func main(){
	desafio1()
	desafio2()
}

func desafio1() {
	for i:= 1; i < 101; i++ {
		if(i%3 == 0){
			fmt.Println(i)
		}
	}
}


/*
"Programa para contar de 1 a 100, se for multiplo de 3, aparece pin, se for multiplo de 5, mostra pan,
se não aparece o numero normal.
*/
func desafio2() {
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