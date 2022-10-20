package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

// implement Preparer interface on type Chicken:
func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (c Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		// make a copy of the current element of the loop:
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	prepareDishes(dishes)
}
