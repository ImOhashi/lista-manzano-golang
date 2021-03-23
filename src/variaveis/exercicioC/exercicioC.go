package main

import "math"

func exercicioC(raio, altura float64) float64 {
	var volume float64 = float64(math.Pi) * math.Pow(2, raio) * altura
	return volume
}
