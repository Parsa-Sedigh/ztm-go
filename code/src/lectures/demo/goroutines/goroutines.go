package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	// a closure:
	/* Every line within this closure, if there's other go routines running, it's possible that they can execute in between each line here. So we have
	to be extra careful when you're working with go concurrency and goroutines.*/
	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}

	// 100 milliseconds should be enough time for the goroutines to capitalize everything and place them into our slice
	/* For the purposes of this demo, we're using time.Sleep() to wait for the goroutines to finish. In most code, there are other mechanisms to wait and
	we'll be covering those in a future video. For now, time.Sleep() will suffice, just make sure we have enough time for the calculations to complete. */
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Captilaized: %c\n", capitalized)
}
