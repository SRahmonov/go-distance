package main

import "fmt"

func main() {

		distance := distance(9.2,8.6)
		fmt.Println(distance)
}
	//volumeFuel -> объем топливо
	//expenseFuel -> расход топливо
	//distance -> расстояние
func distance(volumeFuel float32, expenseFuel float32 )float32{
		carCanRideFor := 100 * expenseFuel / volumeFuel
		return carCanRideFor
	}
