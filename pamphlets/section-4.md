# 04 - Go Programming Types

## 28-001 Structures
Storing data in groups is usually more efficient than storing them on their own, both for running your program and for reading and writing code.

When instantiating a structure, if you don't provide the name of the fields, you need to order them according to the order in their definition of struct. But if you
write their names too, you don't need to comply their order, but you would generally want tot follow the same order as the struct definition itself, that way if you
make updates, it's easier to change it.

So if you don't list the names of the struct when instantiating it, you need to have the fields in order.

Just like with variables, if you try to create a structure and don't supply any data, the fields will have default values.

With anonymous structures, it allows you do define the data structure where you need it instead of having it in a different file or somewhere else across
the source code.

To create anonymous structures, we use the var keyword then the name of the struct and then keyword struct or use create and assign shorthand(:=) and don't
give it a name and just use the struct keyword and define the fields and immediately define the values of them(in this case, you shouldn't miss a value of a field).

So when you use (:=) to create an anonymous struct, you won't be able to have default values for any field of the struct(right side of the pic).

When you use the create and assign shorthand to instantiate an anonymous struct, you won't be able to have default values for any field in the structure.(Look at
anonymous structures example slide). 

## 29-002 Demo Structures
demo>structs

## 30-003 Exercise Structures


## 31-004 Arrays
When we have 3 dots when defining an array:
```go
myArray := [...]int{7, 8, 9}
```
What 3 dots is gonna do is to look over in the curly braces and see how many items you have created and it's gonna substitute that number of items in [] .
So this is kinda default way you'd want to make an array, this way you can easily update the number of items without having to actually worry about what is your number in [].
So 3 dots in this case, is gonna be 3 .

In:
```go
myArray := [4]int{7, 8, 9}
```
the fourth item will have the default value of an `int`, which is 0.

It's good to make a copy of the thing that you're looping through it, into a new variable and this will help because the iteration number(like i, j) will change
in one thread but not the other and if you make a copy, you don't have to worry about any bugs popping up. So it's good to copy the current element in loop(like rooms[i]) into
a new variable.

## 32-005 Demo Arrays
Whenever you're working with the iterator variable(i, j, k) within a for loop in other words, whenever you utilize the iterator variable within a loop, always make a copy, like:
```go
room := rooms[i] // here, we're utilizing the iterator variable with rooms, so make a copy of rooms[i]
```

Even if you use the iterator variable itself without indexing an array, you still wanna make a copy of it. For example, when you want to use the i variable itself,
you can do:
```go
j := i
// now, you can work with j from this point onwards, instead of `i`
```

## 33-006 Exercise Arrays


## 34-007 Slices
In the pic below, you note that 4, 5 and 6 are greyed out, but the 0(in slice) is missing and that's kinda significant. Sp what happens when you create a slice, is if you
**skip** some elements(like the index 0 in pic below), those elements are just **completely** deleted. So you can no longer access those.
However, you're able to access elements that are on the other end of the array(4, 5 and 6) which is what we're doing on the step 3 there.
![img.png](/img/section-44/33-006-1.png)
You can create and destroy slices whenever and change the size and change which part of the array it views and the array always remains untouched because the slice
is simply looking at the original array and not copying it. This means the slice takes a minimal amount of memory, because it just
has an address that points into that existing array. So you're not making a copy when you make a slice, all you're doing is having some metadata that just
peeks into that array.

**To create array and slice at the same time, don't put anything in the []** . Whenever that [] part is empty, it creates an array which will be the size of the length of
the elements and then it creates and gives you a slice which is a view into that array with the length of the elements. For example:
```go
mySlice := []{1, 2, 3}
```

### preallocation:
With preallocation(which means you know ahead of time, how many items that you need), we can save computation time. We can use `make` function to do this.

In function params if nothing is specified in the `[]` of the parameter, it means we can pass an array or slice with any number of elements to it. This is where the
len function becomes important because that's the only way to get the number of elements there.

Slices have no numbers or 3 dots or any symbols in `[]` of them.

Whenever you're working with elements of slice or array inside a loop and you have a iterator variable(or you're working with the iterator variable itself), you wanna make a 
copy of it.


## 35-008 Demo Slices

## 36-009 Exercise Slices

## 37-010 Ranges
When we use the `range` keyword, it's gonna automatically create an iterator for us and so it's kinda like an easeir way of making a for loop with the counter.

We can't access letters of a string by saying:
```go
slice := []string{"Hello"}
for i := 0; i < 10; i++ {
	slice[0][i]
}
```
This is not gonna work, because we're going through individual bytes of that string and the letters take up more than 1 byte in some cases like special
characters or emoji for example. However, with range keyword, it's gonna automatically account for all that for us and so we're able to go through each letter when
we use the range keyword.

If we use %q in print, we'll print out the glyph representation of runes.

With `%q` within printf, we're able to print a rune instead of just the number of the rune.

So the range, for the most parts, simulates the typical `for` loop with the iteration and the counter, except it also takes into account the glyph representation of
these string and runes. one of the most convenient thing about ranges is we don't have to calculate the length of anything(but in for, we typically use len(slice) or ...),
because all of that is considered within the `range` keyword itself.

## 38-011 Maps

## 39-012 Demo Maps
When use `range` on a `map`, it's gonna be a random order.  


## 40-013 Exercise Maps

## 41-014 Pointers
Function calls in go are passed by value. This means that whenever you call a function, a copy of each argument is made, regardless of how large it is.
This means that when you use functions with large data structures such as really big slices or maps with lots of information, it could be potentially slow when you're
calling your functions and this makes it more difficult to manage your program state. Since you're sending copies around everywhere, maybe ine copy has one set of data
and another copy has a different set of data.

We're able to change this default pass by value behavior by using pointers.

`*int` means pointer to integer.

`&value` means pointer to value.

So we use * when we're declaring that we need a pointer to sth and we use & when you wanna actually create a pointer from sth that already exists:
```go
value := 10
var valuePtr  *int
valuePtr = &value
```

The shorter version of above code:
```go
value := 10
valuePtr := &value
```
`valuePtr` is gonna have a memory address and it points to `value` variable.

When we use * with a pointer(and not a type and instead a pointer variable), it will dereference the pointer and what de-referencing does, is it provides
access to the data that's being pointed to.

```go
i := 1
increment(&i)
```
The &i means a pointer to `i`. So with &i we're creating a variable that has a memory address as it's value that points to `i`. This means that increment function
will be able to have access to that `i` variable and it will not have a copy.

## 42-015 Demo Pointers
Whenever we're working with pointers and structures, we know that structures have the dot notation to access fields, now if we have a pointer to a struct, we don't need
to de-reference it with the * . So when we have structs, the dot notation of accessing fields will do that for us automatically and that's because accessing fields
through a pointer with structs is such a common task that it just does that and we don't need to constantly be using asterisks.

## 43-016 Exercise Pointers

## 44-017 Section Review Library
Go to exercises/sr-library (sr is section review).

In order to accomplish this, we need to use pointers! and the reason we only want 1 copy of the library in memory is that's the single source of truth for
what books have been lent out and returned and if we were to make copies of the library, there's a good chance that different copies of the library would have
different information and then it would be hard to keep track of where the books are or where they're supposed to be. So if we just have 1 copy, then all of the
books will go through that single copy.

The first thing I'll be creating, are some type aliases and structures. A type alias example would be:
```go
type Title string
```
which is for the book title.

## 45-018 Exercise Imposter Syndrome

