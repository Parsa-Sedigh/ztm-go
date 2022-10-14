package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	/* With regular variables, we would need to use * to dereference it by saying: *variable += 1 . However when we have structs, the dot notation of
	accessing fields will do that for us automatically. */
	counter.hits += 1
	fmt.Println("Counter", counter)
}

// the new string param gonna be a copy
func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}
	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	// hello variable now should hold "Hi"
	fmt.Println(hello, world)

	/* When we create this phrase variable, we made copies of the existing variables. So hello will have "Hi" */
	phrase := []string{hello, world}
	fmt.Println(phrase)

	/* Here, we change the world that is copied in phrase, so the world variable at the top of the main function is unaffected after this line:*/
	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}
