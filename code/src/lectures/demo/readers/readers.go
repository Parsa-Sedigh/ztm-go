package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input) // get rid of any spaces that are around

		/* If user has entered a bunch of spaces with nothing in between, in that case we don't care what the input is, because there's nothing. So we can
		go on to the next iteration using continue. */
		if n == "" {
			continue
		}

		/* When we get to this point of the program, we have a string and it should be a number of random numbers, so we're gonna try to convert it from a string
		into a number. */
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		/* Normally we would check for errors immediately after running the function that could error. So this check would normally go right after we do r.ReadString() above.
		However, when we get the io.EOF error, that doesn't mean that sth went wrong. That just means that there's nothing left and there could still be input in input variable above.
		So we could have input and we have the end of the file. So we still need to make sure that this input is checked, so we put this error check
		at the end, so that the last input is also checked.

		So this check is not a real error, we're just checking if we run out out data.*/
		if inputErr == io.EOF {
			break
		}

		// an actual error occurred(we were unable to read the user input):
		if inputErr != nil {
			fmt.Println("Error reading stdin:", inputErr)
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
