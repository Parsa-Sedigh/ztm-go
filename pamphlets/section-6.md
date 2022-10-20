# 06 - Interfaces in Go

## 57-001 Interfaces
When we're creating functions, the type of data that's expected by the function must be specified in that function parameters.
However, we don't always know the specific type ahead of time.

When we create an interface, we're saying that here(inside interface definition) is the group of functionality and if that group of 
functionality are present on a type, then this interface is considered to be implemented.

When you're creating functions that use an interface, we should never have a pointer to that interface, it should always just be a value type.
The reason for that is that the caller of the function which receives a param with interface, (the caller) determines whether or not a pointer or a value is
used. So we don't want to restrict the caller to only having a pointer interface because if we do that, they won't be able to use a value.

In **pass by value vs pointer** slide, we know that MyType type alias implements MyInterface interface(because we created those two functions).

Note:
If self-modification is needed, then implement all interface functions as receiver functions for consistency. What this means is if you're going to use a pointer
for one of your implementations, then just use a pointer for all, or if you're gonna use a value type for one of your implementations, then use a value type
for all the functions. Therefore, if you proceed with a particular implementation where one function implementation is a pointer receiver and function2 implementation is a 
value receiver function, the problem lies when we try to call one of the implementation functions, using a copy of value, we're gonna get a compiler error.

Most of the time you will be using pointer implementations, so the above note is not a big deal. But one important thing is to just use all pointer implementation functions
or use all values.

### access implementing type:
Sometimes you need to access the underlying type that implements an interface and you would do this to call specific functions or to make type specific modifications.


### Recap:
Interfaces are implicitly implemented. This means all you need to do to implement an interface is to create receiver functions that match the signatures of the all
the functions within that interface, for your specific type.

When you're writing functions that accept an interface as a function parameter, you don't need to use pointers to interfaces. You usually want to use
pointers at the call side and what this does, is it allows your function to operate on both copies and pointers for that specific interface.

If one of the interface functions that you implement on your type is a pointer receiver function, then the type can only be used as a pointer in function calls.
So usually you wanna either do all pointers, or all value types, that way you don't need to continuously mix and match and get compiler errors later on.

## 58-002 Demo Interfaces
The naming convention for interfaces in go, is to have `er` as a suffix at the end of your interface name.

Interfaces allow you to take multiple types and execute similar functionality on all of them without having to write individual functions for each specific type(your type
just needs to have functions required for implementing the interface, be defined for itself).

## 59-003 Exercise Interfaces

## 60-004 Error Handling
Go has no exceptions. That means there is no try and catch, instead, errors are returned as the last returned value from a function and this encodes the
failure of a function as part of the function signature itself. This makes it simple to determine if a function can fail.

Whenever the function has the possibility of returning an error and the error value is nil, that would mean that no error occurred.

When errors are returned from the functions, they implement the `error` interface that is from the standard library and we only have one function to implement and
that is the `error` function and it returns a string and that string just details what went wrong.

The easiest way to create errors is to use errors.New() , because you don't need to create a type alias or anything, you just use `errors.New()` .

### working with errors:
Since errors is an interface, you don't always know exactly what error occurred. However we can use errors.Is() to determine if the error does have a specific type somewhere
within the errors.

## 61-005 Demo Error Handling

## 62-006 Exercise Error Handling
The testing file needs to be in the same package as our source file that is being tested.

To run the test:
```shell
go test -v ./exercise/errors
```

## 63-007 Readers & Writers
The Read function of Reader interface returns the number of bytes that were read and an error if anything occurred.

Whenever we're creating a buffer, it's setting aside some memory and we're gonna using it as just a temporary location to store some things before we copy it into the
final destination.

EOF = end of file and it means there's nothing to read.

### Walkthrough:
First thing we do is create a reader with "SAMPLE" in it(when we do NewReader() function call).
Currently, the newString has no memory allocated, because it's just an unintialized variable.

buffer variable is 4 bytes in size and right now is empty.

Then we read some data by using `reader.Read(buffer)` into the buffer(we take the first 4 bytes from reader and they get copied into that buffer).`

Then we take a chunk out of that buffer based on the number of bytes: `chunk := buffer[:numBytes]` and that gets copied into newStirng using `Write()` that
is defined on strings.Builder() . 

At this point we're on iteration number 2 and we pull out more bytes from reader into the `buffer` variable and all we have left is `LE`. You see in the pic that
M and P are still greyed out and that's becuase the buffer doesn't erase the old data and that's why we need to have the `numBytes ` and that's important because
if we don't have that, then we're going to have invalid data that's old from the previous iteration and again, the data from buffer is copied into the newString when we use
the write function. On the third iteration we will get an EOF error, our numBytes will be 0, so our chunk will be empty. we won't write anything and we break out of the loop.

By splitting on the new line by using ReadString('\n') , once a new line is reached, then that'll be considered the end of the file.

## 64-008 Demo Readers
After running the program with `go run`, input some data like: 1 1 1 1 . Then you can hit control + d to signal the end of input(on windows it's ctrl+z).

Try running the program with both letters and numbers: 1 a b c 1 1 1

## 65-009 Exercise Readers


## 66-010 Type Embedding
## 67-011 Demo Type Embedding
## 68-012 Exercise Type Embedding
## 69-013 Generics
## 70-014 Demo Generics
## 71-015 Exercise Generics