package worklist

// an entry in the channel:
type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

// it's gonna be a buffered channel, so we have to set the buffer size
func New(bufSize int) Worklist {
	return Worklist{make(chan Entry, bufSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

/* The created empty jobs will be signaling to the workers that it's time for them to quit, meaning that there are no more files available.
So what's gonna happen when we call Finalize is we're gonna take the number of workers and just add a blank entry and whenever a worker pulls a job out
that has a blank entry, that worker is gonna self terminate because there's no more work to be done and that's why we have a for loop in this func, so we need
to add as many blank entries as there are workers, that way every single one will be terminated. */
func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
