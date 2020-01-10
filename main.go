package main

func main() {

}
func distance(fuelConsumed int, fuelInTheTank int) int {
	const distance int = 100
	carCanDriveThrough := distance * fuelInTheTank / fuelConsumed
	return carCanDriveThrough

}