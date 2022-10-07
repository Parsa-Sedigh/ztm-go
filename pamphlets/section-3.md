## 10-001 Variables:
variables are simply an alias to some data in memory and makes it easier to access that data.

![img.png](/img/section-3/10-001-1.png)
![img.png](/img/section-3/10-001-2.png)
![img.png](/img/section-3/10-001-3.png)
![img.png](/img/section-3/10-001-4.png)
![img.png](/img/section-3/10-001-5.png)
![img.png](/img/section-3/10-001-6.png)

`:=` is `create and assign` operator. It creates a copy of the right side to the left side. In example below, if later we change a to 2, b would still be 1. Because
we copied the data that was in `a`.
```go
a := 1
b := a 
```

![img.png](/img/section-3/10-001-7.png)

![img.png](/img/section-3/10-001-8.png)
![img.png](/img/section-3/10-001-9.png)
Why in the pic we can recreate the b variable?
The reason we have this special case, is because lots of code might give you errors and when that happens, we'll have the variable like x assign and possible error values in
the err variable, only to check those errors. But when we use the create and assign operator, it allows us to easily create a new variable **before comma**(it's comma ok idiom!), like x or
y or z, while still **reusing** our err values. If we weren't allowed to do this, we would have to create err1, err2, err3. So it's easier to keep using the same error variable, since it's
going to contain the same type of info.
![img.png](/img/section-3/10-001-10.png)

![img.png](/img/section-3/10-001-11.png)

Note: We don't need to use := but instead = . Also the naming convention for constants is a capital letter for each word even the first word. In C, usually these would be all capital.     


THe comma OK idiom, allows you reuse the second variable(after comma) when you're creating a compound variable and this idiom is frequently used when you're checking for errors.
![img.png](/img/section-3/10-001-12.png)

## 11-002 Demo Variables:

## 12-003 Exercise Variables

## 13-004 Basic Functions:
![img.png](/img/section-3/13-004-1.png)
![img.png](/img/section-3/13-004-2.png)

![img.png](/img/section-3/13-004-3.png)
lhs: left-hand side
rhs: right-hand side

If you have multiple parameters in a function with all of them of the same type, you can just list them with commas and at the very end(for last same type param) indicate which
type they are, just like the above image.


![img.png](/img/section-3/13-004-4.png)
![img.png](/img/section-3/13-004-5.png)
![img.png](/img/section-3/13-004-6.png)
The multiReturn function doesn't have any **parameters**.

![img.png](/img/section-3/13-004-7.png)
![img.png](/img/section-3/13-004-8.png)

Functions are used by calling the function and supplying **arguments** to the **parameters**.
## 14-005 Demo Functions:

## 15-006 Exercise Functions:

## 16-007 Operators:
![img.png](/img/section-3/16-007-1.png)
![img.png](/img/section-3/16-007-2.png)
For example, the += operator takes an existing variable, performs that + operation and **reassigns** the result to that variable again.
Example:
```go
a := 1
a += 3 // shortcut for a = a + 3
```
Here, in the second line, that will take the existing `a` and add 3 to it and then reassign the result to `a` again. So a will be equal to 4.

![img.png](/img/section-3/16-007-3.png)
![img.png](/img/section-3/16-007-4.png)
![img.png](/img/section-3/16-007-5.png)

![img.png](/img/section-3/16-007-6.png)

![img.png](/img/section-3/16-007-7.png)
![img.png](/img/section-3/16-007-8.png)

## 17-008 if..else
![img.png](/img/section-3/17-008-1.png)
![img.png](/img/section-3/17-008-2.png)
![img.png](/img/section-3/17-008-3.png)
![img.png](/img/section-3/17-008-4.png)
![img.png](/img/section-3/17-008-5.png)
![img.png](/img/section-3/17-008-6.png)
![img.png](/img/section-3/17-008-7.png)
Remember with statement initialization, we need to have the if keyword, followed by our initialization statement, a semicolon and then our comparison at the end.
The comparison should always refer to the variable that you created in the initialization statement. When you do so, you'll be able to access that created variable(in init statement) within
the blocks.

![img.png](/img/section-3/17-008-8.png)
An early return is simply using the if statement to validate some data. When the data is validated, you're able to continue with the function. If the dat fails validation,
you return early from the function.
![img.png](/img/section-3/17-008-9.png)
If both of the err variables are nil(so we don't return early at all), we will end up down at the last if and we'll have valid token data and 
valid cart data. In the bodies of `if err !== nil {}`, you can perform logging if you want. 

![img.png](/img/section-3/17-008-10.png)

## 18-009 Demo if..else

## 19-010 Exercise if..else

## 20-011 switch:
In this vid:
- basic usage
- case lists
- fallthrough behavior

It's always a good practice to have a default case and that's just in the event that you forget to check for a case, at least we have sth that will be executed.
![img.png](/img/section-3/20-011-1.png)
![img.png](/img/section-3/20-011-2.png)

![img.png](/img/section-3/20-011-3.png)

Case list allows us to check for multiple things in the same case. Case lists are useful if you want to take the same action for multiple different values.
![img.png](/img/section-3/20-011-4.png)

### fallthrough:
The default behavior for switches in go is to automatically stop executing once one of the cases matches.
But with fallthrough, if a case matches and it has a fallthrough keyword there, the next case will be automatically executed even if that next case doesn't match with the value
we're switching, For example with the 'i' in the next example, the second case will be executed and because of fallthrough, the next case will also be executed even though it doesn't
match with 'i'.
![img.png](/img/section-3/20-011-5.png)

The fallthrough behavior can be useful if you want to always execute some code, regardless of which case was matched.
![img.png](/img/section-3/20-011-6.png)
## 21-012 Demo switch:

## 22-013 Exercise switch

## 23-014 Looping
- basic loop
- while loop
- infinite loop

![img.png](/img/section-3/23-014-1.png)
![img.png](/img/section-3/23-014-2.png)
![img.png](/img/section-3/23-014-3.png)

TODO: Till start of infinite loop
Infinite loops are used in web servers or if if you're reading data from some kind of data stream and then you would choose a condition on which to break.

![img.png](/img/section-3/23-014-4.png)
![img.png](/img/section-3/23-014-5.png)

As soon as we hit the continue, we will go up to the for loop again and the condition will be checked again to figure out do the same iteration or not. So the code after
continue(inside the if or outside of if after it) will never get executed if the continue is hit first.
![img.png](/img/section-3/23-014-6.png)
![img.png](/img/section-3/23-014-7.png)

## 24-015 Demo Looping
In:
```go
for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}
```
Whenever the loop body executes and gets the end, the post statement will be run(in this case i will be incremented by 1) and the condition will be checked again.
Run the demo with:
```shell
go run ./demo/loops
```

## 25-016 Exercise Looping

## 26-017 Section Review Dice Roller
Go package registry: [link](https://pkg.go.dev/std)

Go to sr-dice exercise.

We're gonna need random numbers for dice rolling.

The first thing we need to od is seed the random number generator. What seeding does is it sets the initial value for the random numbers. So if you use the
same seed value, whenever you run the program, you will get the same random numbers. To get random numbers in each program run, use the current time in `rand.Seed()`. So the
seed value is gonna be the current time which is gonna be different between program runs.

The UnixNano function will get the number of nanoseconds since the start of the unix epoch.

## 27-018 LinkedIn Endorsements