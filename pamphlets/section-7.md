# 07 - Concurrent Programming with Go

## 72-001 Function Literals
Function literals provide a way to define a function within a function. This makes it possible to assign function literals to variables.

Function literals are also known as closures or anonymous functions in other programming languages and closures also allow data to be encapsulated within.

In addition to just being a function literal, a closure will also take values from outside of the functions.

Whenever you have sth outside of your function literal and you use it within the function, you're gonna get a copy of that within your function.

What closures allow us to do is have our own functionality in small chunks, send this functionality to different functions and then run it within that one and this is
how you can make your code do different things in different situations and also makes it easy to add new functionality to existing code.

## 73-002 Demo Function Literals
## 74-003 Exercise Function Literalas
## 75-004 Defer
With defer, each of the function calls that have defer on them, is gonna be placed on a stack. So they're gonna be executed in reverse order.
The last defer will be executed first and ... (they're gonna be executed backwards).

In `` slide, the reason we placed that defer file.Close() there, is because at that point in the program, we do have a handle to a file and you always want to
close it after you open it. That way we don't have resources dangling around.

## 76-005 Concurrent Programming
When we have concurrency, we're allowed to execute multiple different lines at a time.

Async code: code that can pause and resume when needed. When async code is paused, other async code can resume. So usually you only have 1 async piece of code
running at one time and while it pauses, other pieces will be running.

One of the great things about go is it automatically chooses threaded or asynchronous, so we don't have to worry about the details on how to manage threads
or async code because go handles those things for us.

### Threaded execution slide:
There we have a CPU core(the big one) and it executes code on the main thread. We then have 4 other cores that can execute code and the main thread can branch off
into these cores and those cores can do their own thing and then join their data back onto the main thread. So when we get the data back, we're joining on the
main thread, so that grey line can be a thread and then once we get data back, we join on the main one and generally you wanna wait for all these jobs to finish. The
main thread would just pause right after that last grey line joining and once all those cores are done, all the results will go in the point where
last grey line is joined with the main thread and once all the results are in, then execution of the main thread would continue.

In any kind of concurrent programming, you do have to wait on your main thread, before you can continue on with the program, because if you just continue on with your
program, it could exit and you can still have cores trying to compute results and if your main program exits, then all of these jobs are gonna get terminated and so
you won't have the correct results.

All concurrent code whether it's threaded or async, runs non-deterministically. This means that each time you run the program that uses concurrency, you'll get
results in a different order.


## 77-006 Goroutines

## 78-007 Demo Goroutines
When we launch a goroutine using `go` keyword, main thread just keeps on going and your goroutine is doing it's own thing in the background and so you have
2 separate threads of execution going at the same time.

Closure captures are shared among all your goroutines, this makes it easy to parallelize code.

To run a program multiple times(like 15 times):
```shell
for i in {1..15}; do go run <path to the go file> | rg ':'; done
```

Always take extra care when you're writing multi-threaded code using goroutines.

## 79-008 Exercise Goroutines
Sine this exercise has additional data files, we need to change directory. That way our program could read it. So:
```shell
cd excercise/goroutines
go run .
```
In this exercise, since we're doing a sum operation by just adding things together, the order of the operations don't matter, so we can have the goroutines running
in any order and it's a non-issue. However, if we had a different operation, such as division, we would have to make sure that we have the goroutines executing in
the proper order. This will be accomplished by using different techniques.

## 80-009 Channels
We're not able to pull out messages in the middle of the channel, so it's just one goes in and then one comes out.

### multiple receive ends:
With multiple receive ends, we can have one goroutine sending data into a channel and we can have for example 3 separate goroutines, all reading data from that
channel. This allows us to easily split up our work across multiple goroutines.

### Creating and usage:
We may not get the results back in the same order as the previous run if we're utilizing goroutines, it could be in any order. Because the goroutines
can place the data into the channel in a non-deterministic order.

Unbuffered channels will **block** when sending, until a reader is available. This means that when your line of code executes and sends data nnto a channel,
it will sit there on that line of code waiting until a reader in a different goroutine reads that data. Buffered channels on the other hand, have a specific capacity
and you can send messages into the channel up to that capacity, even if there is no reader on the other end pulling data out.

Channels use FIFO ordering: The first message you place into the channel, will be the first one taken out of the channel.

We're not spawning a goroutine for 1 and 2, because it's buffered and so the execution will just keep continuing until the channel fills up and the channel fills up
with 2 numbers. Now in the anonymous function, we add the third message onto the channel and this time we're inside a goroutine. That's because we only have a
buffer capacity(capacity of the channel) of two and our program will block on the third attempt to add data. So instead of just blocking the main program,
we create a new goroutine in that anonymous function and this goroutine within that {}, that's gonna sit there blocking until there's a room available in the channel.
Since we already added two messages, we're unable to add anymore and that anonymous function is gonna sit there doing it's own thing as a goroutine and since it's a gorutine,
we can continue on with our **main** program(but that anonymous function will be blocked).

One difference between this program(in `` slide) and other programs, is since we're running deterministically at this point:
```go
channel <- 1
channel <- 2
```
The above lines are running on the main thread which is running deterministically.
we're gonna get the results out as 1 then 2 and then 3. Why 3 at the end?
Because that goroutine on that anonymous function is only one goroutine.

Generally you'll have multiple goroutines running at the same time, so it's gonna be non-deterministic most of the time. However, the tutor just wanted
to show that you can utilize channels deterministically by having them as part of the main thread or within a single goroutine.

## 81-010 Demo Channels
## 82-011 Exercise Channels
## 83-012 Synchronization
### Mutex
In that slide which has a pic, this is some explanation:

We have some shared data which is in the middle of pic and we have goroutine A and goroutine B. Execution will be going downwards. The first thing that happens is
goroutine A will attempt to lock the data. Since the data has no lock right now, the lock will be acquired by goroutine A and then goroutine A has
**exclusive access to the data** until it is unlocked and the unlock operation has to be done by goroutine A.

Once goroutine A unlocks data, it then becomes shared again and then goroutine B is able to acquire a lock.

While goroutine A had lock acquired, goroutine B also attempted to get a lock, however we can see that goroutine B was blocked(that red text) and that's because
the mutex was already acquired by goroutine A. While goroutine B was blocked, it's unable to execute. So those purpule dashed lines indicate that the gorutine B is
just waiting until it can acquire lock and it's able to do so further down, because goroutine A unlocked it at that point. Once B acquires lock, then
goroutine B can do whatever it wants with the data, then unlocks it, the data is shared again where any goroutine can access it and ... .

### Deferred unlock:
Since the defer keyword allows us to execute code when a function finishes, if we just defer the unlock operation, that means that will always make sure to unlock
the locks no matter what happens within our functions.

When you're working with mutex, it's important that they're always unlocked because a locked mutex will freeze your program if it's not properly unlocked. So
you should almost always use `defer` **immediately** after taking your mutex lock. That way, if anything happens later on in your function, you don't have
to worry about it, the mutex will alawys be unlocked and other goroutines will have access to the data.

So far in the concurrency section, we've mostly just been waiting for the job to finish by setting a specific timeout and hoping it just finished within that time.
Waitgroups give us much more control and allow us to wait until the goroutines are actually done with this job before continuing.

What you wanna do with the waitgroups is every time you open a goroutine, you want to call the `Add()` function and add 1 , because you're making 1 goroutine and then
within the function that has `go`, you just wanna call `Done()` when that function is finished and then Done function on waitgroup is gonna reduce that counter by 1.

In the example, we used defer with wg.Done() , this way if anything happens later in the code, I don't have to worry about it, the waitgroup counter will always go down because
of defer execution when function finishes.

The wg.Wait() is gonna pause the execution of the code on that line, until the counter reaches zero again.

## 84-013 Demo WaitGroups
When you call `wg.Wait()` , if the wg counter is greater than 0, then that code will block on the line of Wait() , until that counter reaches 0 and once that occures,
execution will continue, which means all of our goroutines are done.

Once you input your data, you can hit control + d on linux or control + z on windows to send end of file signal.
In mac you can hit control + d twice to send this signal.
## 85-014 Demo Mutexes
## 86-015 Exercise Synchronization
## 87-016 Section Review Multithreaded grep
worklist folder is for work processing queue.

In the mgrep directory, run:
```shell
go mod init mgrep
```

What worklist package does is it keeps track of all the files that need to be processed. So we're gonna create a channel and a few structures.

When a match is found, we're gonna output 3 things(Result struct):
- full line of text where match is found
- indicate which line number in the file that the match is found at
- path of the file that we're working with

You can run the program on the current directory by running this in the root directory of project:
```shell
go run ./mgrep func . # to get every func in the texts
```
The first time you run the program, you'll probably get an error message and that's because we need to download the 3rd party package that we're using.
To download that, run this on the directory that contains `go.mod`:
```shell
go mod tidy
```

Whenever you see sth like: `go get <url of package>`, it's best to instead run: `go mod tidy` just in case you have other modules to work with as well.

