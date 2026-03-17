package elon

import "fmt"

// Drive drives car one time
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance returns the distance driven so far as a formatted string
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery status as a formatted string
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish returns if the car can finish the track
func (car Car) CanFinish(trackDistance int) bool {
	return car.battery >= trackDistance/car.speed*car.batteryDrain
}
