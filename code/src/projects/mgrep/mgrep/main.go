package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"
)

// locate all the directories and file(discover directories).
// the path arg is gonna be the initial search path
func discoverDirs(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}

	/* once all the files and directories have been retrieved, we're gonna exit out of this for loop and return from the function. At that point,
	our worklist is gonna have a list of all the paths that are files for our goroutines to be checking.*/
	for _, entry := range entries {
		/* If we do run into a directory, then we need to recurse into it: */
		if entry.IsDir() {
			/* For example if our initialPath was a and we found a directory called b, then filepath.Join() will return a/b , that way we will
			be able to enter into the b directory without issues.*/
			nextPath := filepath.Join(path, entry.Name())

			// recurse into that directory
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

/* To utilize the go-args package, we're gonna create an args struct.
We don't have to provide a SearchDir and in that case, we should just default to having the current directory be searched, but the SearchTerm is required.
So the program should abort if we don't supply any term to search for.*/
var args struct {
	SearchTerm string `arg:"positional,required"`
	SearchDir  string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	// create a wait group for our workers:
	var workersWg sync.WaitGroup

	// we can have 100 jobs before things start to block
	wl := worklist.New(100)

	// results is where workers put the results in:
	results := make(chan worker.Result, 100)

	numWorkers := 10

	// we're gonna start a goroutine at the next lines, so add 1 on the workersWg
	workersWg.Add(1)

	/* This goroutine is gonna be the one that discovers directories. So we're only gonna have 1 goroutine that goes to the filesystem, pullint out files,
	it's gonna add them to the worklist and as they're being added, our worker is gonna be pulling them out of the worklist and processing them.*/
	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	// spawn our workers:
	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()

			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if workerResult != nil {
						for _, r := range workerResult.Inner {
							results <- r
						}
					}
				} else {
					/* path is blank, so this worker should terminate. We can just return, since we deferred workersWg.Done() , we don't have to worry about that, since
					workersWg.Done() gonna be executed, our counter for wait group will decrement.*/
					return
				}
			}

		}()
	}

	/* We need to wait on the workers to finish. However, we also want to display results while they're working. We're gonna be using the worker's wait group's Wait function,
	however, that's blocking and we want to be able to display results while we wait for the workers to finish. To make that work, I'm gonna wrap it up into a
	channel.

	We can just make a channel with an empty structure and because all we need to do is read it.

	Here, we have a new channel and a new goroutine. That goroutine will just sit there blocking, waiting for the workers to finish and later on,
	we're gonna be selecting on blockWorkersWg channel, because we don't want to block, because we also want to be displaying results at the same time.
	So we combine the two. We can have sth that blocks the execution in it's own goroutine and then combined with the channel, we can also select
	another things. So that allows us to have multiple things going at once, even though some of it is blocking.*/
	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()

		/* This will indicate to the rest of the program that there's no more results coming in and all that's left to do is print'em out. */
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup

	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <-results:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)

				/* When we close the channel, this case always succeeds and since this will always succeed once the channel is closed, we need to
				check the length of the results, because it's possible that our results are not done printing out yet. So once we ran out of results by saying:
				len(results) == 0, we can call Done on the wait group.*/
			case <-blockWorkersWg:
				/* If length is greater than 0, we're not gonna do anything and we'll go back up to the select again and we'll try to get more
				results out and the reason we need to do that is because it's possible that the workers are all finish, but we're still trying to print out
				results, cause workers are running really quickly in the background*. However, printing to the terminal is kinda slow, so it's definitely a scenario
				where the workers are done and we're just waiting for data to get streamed to the terminal. */
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()

	displayWg.Wait()
}
