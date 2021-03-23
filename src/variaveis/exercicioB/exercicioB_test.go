package main

import "testing"

func TestExercicioB(t *testing.T) {
	resultado := exercicioB(59)
	if resultado != 15 {
		t.Errorf("Resultado esperado %d, obtida %f", 15, resultado)
	}
}
