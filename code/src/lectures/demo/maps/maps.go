package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	// by default, if there is no key with that name in that map and the value is an int, it will be 9 and here, we just increment it
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")

	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"])

	/* in the event that `found` is false, cereal will have a default value(in this case, 0). So in case you forgot to check with found,
	you can check the cereal whether or not it has a default value.*/
	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList { /* Everytime this loop runs, you might get a different orders. */
		totalItems += amount
	}

	fmt.Println("There are", totalItems, "on the shopping list")
}
