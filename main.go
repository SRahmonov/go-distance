package main

func main() {

}

func distance(volumeFuel int, consumptionFuel int) int {
	const maxDistance = 100
	carCanRideFor := maxDistance * consumptionFuel / volumeFuel
	return carCanRideFor
}
