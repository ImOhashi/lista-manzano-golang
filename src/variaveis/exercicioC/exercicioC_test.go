package main

import (
	"testing"
)

func TestExercicioC(t *testing.T) {
	resultado := exercicioC(1, 2)
	if resultado != 12.566370614359172 {
		t.Errorf("Resultado esperado %f, obtida %f", 12.566370614359172, resultado)
	}
}
