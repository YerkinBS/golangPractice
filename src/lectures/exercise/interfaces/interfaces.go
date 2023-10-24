//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Handler interface {
	HandleVehicle()
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) HandleVehicle() {
	fmt.Println("small lift")
}

func (c Car) HandleVehicle() {
	fmt.Println("standart lift")
}

func (t Truck) HandleVehicle() {
	fmt.Println("large lift")
}

func HandleVehicles(vehicles []Handler) {
	fmt.Println("Handling vehicles:")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		fmt.Printf(" --Vehicle: %v --> ", vehicle)
		vehicle.HandleVehicle()
	}
	fmt.Println()
}

func main() {
	vehicles := []Handler{Motorcycle("Yamaha"), Car("Toyota"), Truck("Zeekr")}
	HandleVehicles(vehicles)
}
