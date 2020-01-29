package main

func main() {

}

func distance(volumeFuel int, consumptionFuel int) int {
	const distanceInKm = 100
	carCanRideFor := distanceInKm * consumptionFuel / volumeFuel
	return carCanRideFor
}
