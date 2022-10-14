package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	// indicates who is on the front seat of the bus?
	FrontSeat Passenger
}

func main() {
	/* one line instantiate technique and we have to list out all the fields in this technique(because we didn't specify name of the fields, we have to
	write them in order) */
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	/* create structures using var keyword(I list out name of the fields, still it's better to write them in order). Because we didn't use the := , we don't need to
	specify all the fields.*/
	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	// here, the fields of heidi have default values

	// now they would have values in these 2 lines:
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println("Casey has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
