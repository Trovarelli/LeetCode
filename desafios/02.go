package desafios

import (
	"fmt"
	"math"
)

func solution02(partida [2]float64, destino [2]float64) float64 {
	minY := math.Abs(partida[1] - destino[1])
	minX := math.Abs(partida[0] - destino[0])

	if math.Ceil(partida[0]) == math.Ceil(destino[0]) {
		route1 := (math.Ceil(partida[0]) - partida[0]) + (math.Ceil(destino[0]) - destino[0])
		route2 := (partida[0] - math.Floor(partida[0])) + (destino[0] - math.Floor(destino[0]))
		minX = math.Min(route1, route2)
	}

	if math.Ceil(partida[1]) == math.Ceil(destino[1]) {
		route1 := (math.Ceil(partida[1]) - partida[1]) + (math.Ceil(destino[1]) - destino[1])
		route2 := (partida[1] - math.Floor(partida[1])) + (destino[1] - math.Floor(destino[1]))
		minY = math.Min(route1, route2)
	}

	return minX + minY
}

func Desafio02() {
	partida := [2]float64{0.4, 1}
	destino := [2]float64{0.9, 3}
	fmt.Printf("O caminho mais curto entre as coordenadas [0.4, 1] e [0.9 e 3] é: %f", solution02(partida, destino))
}
