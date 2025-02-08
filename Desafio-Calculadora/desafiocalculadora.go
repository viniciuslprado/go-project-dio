package main

import "fmt"

func main(){
	x := subtracao(10,5)
	y := soma(10,5)
	z := divisao(100,5,2)
	w := multiplicacao(10,5)
	
	fmt.Println(x,y,z, w)
}

func subtracao( i... int) int{
	total:= 0
	for _, v := range i {
		total = total - v
	}
	return total
}

func soma( i... int) int{
	total:= 0
	for _, v := range i {
		total += v
	}
	return total
}

func divisao( i ... int ) int{
	total := i[0]
	for _, v := range i[1:] {
		total /= v
	}
	return total

}

func multiplicacao( i ... int ) int{
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total

}