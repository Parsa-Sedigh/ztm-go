package main

import "fmt"

// first let's create some constants and these indicate the automated conveyor belt(they can go on to small, medium or large conveyor belt)
const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

// we're gonna use iota enumeration pattern on the Shipping and BeltSize. This way we can get some strings.
func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

// let's create receiver functions to implement the interfaces:
// we can automate SpamMail:
func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

// create sth that can not be automate it's shipping:
type ToxicWaste struct {
	weight int
}

// we need to ship ToxicWaste via ground
func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	/* We're unable to use waste with automate function because we didn't implement Shipper interface for ToxicWaste, because it's dangerous! */
	//waste := ToxicWaste{300}
	//automate(&waste)
}
