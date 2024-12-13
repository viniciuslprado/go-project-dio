/*"Programa para mostrar todos multiplos de 3 de 1 a 100."*/
package main

import "fmt"

func main() {
	for i:= 1; i < 101; i++ {
		if(i%3 == 0){
			fmt.Println(i)
		}
	}
}