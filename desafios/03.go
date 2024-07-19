package desafios

import "fmt"

type Desafio03T struct{}

func Desafio03() {
	carDimension := [2]int{3, 2}
	parkingLot := [4][6]int{
		{1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1},
	}
	luckySpot := [4]int{1, 1, 2, 3}

	fmt.Println(solution03(carDimension, parkingLot, luckySpot))
}

func solution03(carDimension [2]int, parkingLot [4][6]int, luckySpot [4]int) bool {
	luckySpotIsEmpty := isEmpty(parkingLot, [2]int(luckySpot[:2]), [2]int(luckySpot[2:]))

	if !luckySpotIsEmpty {
		return false
	}

	if luckySpot[3]-luckySpot[1] > luckySpot[2]-luckySpot[0] {
		// ENTRADA HORIZONTAL
		return isEmpty(parkingLot, [2]int{luckySpot[0], 0}, [2]int{luckySpot[2], 1}) ||
			isEmpty(parkingLot, [2]int{luckySpot[0], luckySpot[3]}, [2]int{luckySpot[2], len(parkingLot[0]) - 1})
	} else {
		// ENTRADA VERTICAL
		return isEmpty(parkingLot, [2]int{0, luckySpot[0]}, [2]int{luckySpot[0], luckySpot[3]}) ||
			isEmpty(parkingLot, [2]int{luckySpot[2], luckySpot[1]}, [2]int{len(parkingLot) - 1, luckySpot[3]})
	}
}

func isEmpty(parkingLot [4][6]int, p1, p2 [2]int) bool {
	for i := 0; i < len(parkingLot); i++ {
		if i < p1[0] || i > p2[0] {
			continue
		}
		for j := 0; j < len(parkingLot[i]); j++ {
			if j < p1[1] || j > p2[1] {
				continue
			}

			if parkingLot[i][j] == 1 {
				return false
			}
		}
	}

	return true
}
