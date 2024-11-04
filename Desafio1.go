/*"Programa para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius."*/
package main

import "fmt"

const ebulicaoK = 373.2 

func main() {
	tempk := ebulicaoK
	tempc := tempk - 273.2
	fmt.Printf("A temperatura da ebulição da agua em Kelvin é %g, e em Celsius é %g .", tempk, tempc)
}
