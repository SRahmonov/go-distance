package main

func main() {

}
	//volumeFuel -> объем топливо
	//expenseFuel -> расход топливо
	//distance -> расстояние

func distance(volumeFuel float32, expenseFuel float32 )float32{
		carCanRideFor := 100 * expenseFuel / volumeFuel
		return carCanRideFor
	}
