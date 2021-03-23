package main

import "testing"

func TestExercicioA(t *testing.T) {
	resultado := exercicioA(15)
	if resultado != 59 {
		t.Errorf("Resultado esperado %d, obtida %f", 59, resultado)
	}
}
