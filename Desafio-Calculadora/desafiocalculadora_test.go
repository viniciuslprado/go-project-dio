package main

import "testing"

func ShouldSumCorrect(t *testing.T){
	teste := soma(1,2,3)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldMultCorrect(t *testing.T){
	teste := multiplicacao(1,4)
	resultado := 4

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldDivideCorrect(t *testing.T){
	teste := divisao(20,2)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldSubtractCorrect(t *testing.T){
	teste := subtracao(3,2)
	resultado := 1

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldSumIncorrect(t *testing.T){
	teste := soma(1,2,3)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldMultIncorrect(t *testing.T){
	teste := multiplicacao(1,4)
	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldDivideIncorrect(t *testing.T){
	teste := divisao(20,2)
	resultado := 15

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}

func ShouldSubtractIncorrect(t *testing.T){
	teste := subtracao(3,2)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornao", teste)
	}
}


