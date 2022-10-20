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

// instead of using 1, 2, 3 in our code(numbers), we use constants for better readability and also gonna create a type alias
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

// this interface is gonna pick which lift a vehicle should be going to?
type LiftPicker interface {
	PickLift() Lift
}

// the string name for Motorcycle, Car and Truck would be the model name
type Motorcycle string
type Car string
type Truck string

// vehicles have a model name, in addition to the vehicle make:
func (m Motorcycle) String() string {
	/* We know m is a string so maybe it's not required to call string() on m. BUUUT, when you use Sprintf and you have a string type, it's going to
	call this String method as well, so what's gonna happen, is you'll have an infinite loop. */
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

// implement the LiftPicker interface for each one of the types of vehicles:
func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

/* Since we implemented the LiftPicker interface on our Motorcycle, Car and Truck, we will be able to use the sendToLift function with all three of those
different types.*/
func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
