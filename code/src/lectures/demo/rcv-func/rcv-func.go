package main

import "fmt"

// represents a single parking space
type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// We used a pointer instead of copying the value, to change that passed variable itself, not it's copy.
func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial", lot)

	// these 2 functions are the same:
	lot.occupySpace(1)
	occupySpace(&lot, 2)

	fmt.Println("After occupied: ", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate: ", lot)
}
