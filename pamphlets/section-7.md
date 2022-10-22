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
## 81-010 Demo Channels
## 82-011 Exercise Channels
## 83-012 Synchronization
## 84-013 Demo WaitGroups
## 85-014 Demo Mutexes
## 86-015 Exercise Synchronization
## 87-016 Section Review Multithreaded grep


