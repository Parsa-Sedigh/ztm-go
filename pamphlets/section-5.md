## 46-001 Receiver Functions
Similar to modifying a class variable in other langs, for example:
`class.field = <value>`

Most of the time when you create receive functions, they'll probably be pointer receivers. However, there are some use cases for value receivers
such as when you have a small amount of data or if you're working with a slice or a map, since those are **already** pointers by themselves.

## 47-002 Demo Receiver Functions
## 48-003 Exercise Receiver Functions
We use function receivers to make sure that we can edit player stats, you can use normal functions as well, but yiu have to pass a pointer to that struct to
normal function too.

## 49-004 iota
Iota is used when you're working with constants. When we're working with constants, it's common to group them together to represent different statuses or states.

if you're making lots of constants in a group, it's easy to just use the iota keyword instead of assigning numbers to those constants. With iota, the compiler is gonna
take care of filling out all these numbers for you.

### what is iota keyword?
### How to set custom values?
### iota enumeration pattern
The second slice of iota enumartion pattern is a shorter version of the previous slide and it accomplishes this shorter code by using a slice instead of switch.
In this second version, the order in slice should be the same as the constants defined with `iota`.

## 50-005 Exercise iota
## 51-006 Variadics
## 52-008 Packages
## 53-009 Init Function
## 54-010 Testing
## 55-011 Demo Testing
## 56-012 Exercise Testing