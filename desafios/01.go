package desafios

import "fmt"

func Desafio01() {
	ride_time := 30
	ride_distance := 7
	cost_per_minute := [4]float32{0.2, 0.35, 0.45}
	cost_per_mile := [4]float32{1.1, 1.8, 2.3, 3.5}

	var res []float32

	for i, cMinute := range cost_per_minute {
		cMile := cost_per_mile[i]
		calculedRes := (cMinute * float32(ride_time)) + (cMile * float32(ride_distance))
		res = append(res, calculedRes)
	}

	fmt.Println(res)
}
