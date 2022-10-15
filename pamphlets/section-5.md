# 05 - Idiomatic Go

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
The second slice of iota enumeration pattern is a shorter version of the previous slide and it accomplishes this shorter code by using a slice instead of switch.
In this second version, the order in slice should be the same as the constants defined with `iota`.

## 50-005 Exercise iota

## 51-006 Variadics
There will be situations where you write a function and you won't know how many parameters that you need ahead of time. So it's possible to write a function
that accepts any number of parameters nad that's where variadics comes in. So variadtics are a way to make a function that accepts any number of parameters.

## 52-008 Packages
Go to pkg folder in demo folder.

When we create packages, we need to create folders for each one. So create a new folder in pkg folder named `display`, msg and name.

The `package main` is the one that launches the program.

Public and private in go is based on whether or not you capitalize the first letter of a function or structure or type alias or ... .

Run the program by executing:
```shell
go run ./demo/pkg/main 
```

## 53-009 Init Function
When your other packages run their init function, even if you've imported them multiple times, the `init` function will only run a single time.
## 54-010 Testing
In the function which tests another function, we name it Test<name of the function we're testing it>.

The string we pass to Errorf, usually is this pattern: `"<name of the fucntion that is under test>(%v)=<answer you got>, want <answer that you actually want>"`.
For example:
```go
t.Errorf("IsValidEmail(%v)=false, want true", data)
```

In the above string, %v is the data you supplied to the function under test.

Normally you want to use Errorf because it helps to know exactly what failed and why?

The behavior of `Fail` and `Errorf` are to continue trying to run tests, even if one fails, but `FailNow` and `Fatalf` won't continue the tests.

Test tables are like parameterized testing in other languages. With test tables, you specify a set of data and the proper result that we should be getting from
the function. Like:
`table := []struct {
    email string
    want bool
}{
    {"email@example.com", true},
    {"missing@tld", false},
}`

With `-v` flag of `go test` command, it displays all the things that are being tested, otherwise, it just indicates success or failure.

Run:
```shell
go test -v ./demo/testing
```

## 55-011 Demo Testing
## 56-012 Exercise Testing