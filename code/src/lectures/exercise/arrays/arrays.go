//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price int
	name  string
}

func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i] // copy the item to avoid bugs in
		cost += item.price

		/* Why we're doing this? Since our function accepts an array that has 4 products. However one of our requireemtns is we need to have 3 products
		first and then add a product into this array and in order to find out if the product is in the list or not, we should check for the default
		value of string which is an empty string.*/
		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list: ", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost", cost)
}

func main() {
	shoppingList := [4]Product{
		{1, "Banana"},
		{6, "Meat"},
		{3, "Salad"},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{4, "Bread"}

	printStats(shoppingList)
}
