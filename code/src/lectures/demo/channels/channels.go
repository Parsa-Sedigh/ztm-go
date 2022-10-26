package main

import (
	"fmt"
	"time"
)

/* We'll create a program that has a goroutine and will send data through one channel and will have a control channel to control the behavior of the
goroutine */

type ControlMsg int

/* We will have 1 channel that has data for processing(data of type int) and the other channel will return results and they both can use the same structure which
is this.

Note: The data field is the actual job and the second field is a result.*/
type Job struct {
	data   int
	result int
}

const (
	/* DoExit message will be from our main thread to our goroutine and that'll signal it to exit */
	DoExit = iota

	/* ExitOk will be the return value from the goroutine indicating to the main thread that the goroutine has terminated */
	ExitOk
)

func doubler(jobs, results chan Job, control chan ControlMsg) {
	for {
		// select a channel
		select {
		// if we get a message from the control channel(it's gonna uses either DoExit or ExitOk), then we're gonna switch on that message:
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				/* send a message back on the control channel, that we have successfully exited and then the main thread is gonna be reading that channel, just to
				make sure that this goroutine doesn't de-terminate. */
				control <- ExitOk
				return
			default:
				panic("unhandled control message")
			}

		case job := <-jobs:
			results <- Job{data: job.data, result: job.data * 2}
		default:
			time.Sleep(50 * time.Millisecond)

		}
	}
}

func main() {
	// create a buffered channel and the communication is gonna be the `Job` structure and we can put 50 jobs before our main thread will block.
	jobs := make(chan Job, 50)

	// the results channel will be used to control the goroutine itself
	results := make(chan Job, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	// create some jobs:
	/* Here we create 30 jobs and the goroutine is gonna start to chugging away on these jobs immediately and placing results back on the results channel which
	we can get in the select down here.*/
	for i := 0; i < 30; i++ {
		jobs <- Job{i, 0}
	}

	for {
		/* If the first case is got(if we do get sth out of the results channel), the second case is gonna get ignored, we're gonna back to the select(since
		we're inside an infinite for loop) and then it's gonna repeat the process. This means that even if 500 milliseconds elapses, if there are still results
		left in the results channel, we're gonna go to the first case here regardless. So even if 10 seconds elapses, we're still gonna see all the
		results before the loop goes to the first again and then, when there is nothing in the results chan, the second case is gonna executed and that could only
		happen when we have an empty result because the first case always gets checked first(we get results from a chan, one at a time).

		So this structure will allow us to get all the results, no matter what, even if it goes after the amount of time that is allowed for it to complete(500 milliseconds)*/
		select {
		case result := <-results:
			fmt.Println(result)

			// goroutine is gonna have 500 milliseconds tao calculate everything and if elapses, send a control message that we want to exit
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			// if we got this, it means the goroutine got exited
			control <- DoExit
			// wait on a response from the goroutine(ignore the sent value)
			<-control
			fmt.Println("program exit")

			// return from the main function:
			return
		}
	}
}
