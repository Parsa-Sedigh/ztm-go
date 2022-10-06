## 1-

# Section 02 - Introduction to Go (Golang)

## 2-001 So You Want To Learn Go

## 3-002 Install Golang

## 4-003 Packages & Modules:

## 5-004 Monthly Coding Challenges, Free Resources and Guides

## 6-005 Data types
- all primitive data types in go are numeric. This means that they're simply a stream of bytes.
- it's possible that the data is invalid for the given type. this only applies when working with user input or if you're manually manipulating the binary data itself.
  so for the most part, the type that's indicated in the code is the correct type

Unsigned integers are integers, but they;re positive only and they include 0.

uintptr is the pointer size under current architecture. What this means, on 64bit systems, uintptr type will be uint64, on 32bit systems, it'll
be a uint32 and ... .

float64 is a decimal point number with enhanced accuracy.

type alias assigns a new name to an existing type. Useful for providing indication of what kind of data is being utilized.

## 7-006 Strings  Runes:
- encoding is a way to represent thousands of different symbols using code pages
- code pages are tables which use the first few bytes of data to determine which page to use
- each symbol in the code page is called a code point

In the pic, for `d`, we see that we're at code point 1(horizontal) and 3(vertical). **So to create a `d`, it actually takes 2 bytes**. We have to have a 1 and a 
Another example: If we want to create that japanned character, we have to use `4f` followed by `1`. So that particular character takes 3 bytes: `4, f, 1`.

![img.png](/img/section-2/7-006-2.png)

In go, when we work with strings in runes, we're actually working with individual bytes, not letters themselves. 
**So it's important to understand that letters can consume more than one byte.**

![img.png](/img/section-2/7-006-3.png)

- Text in go is represented using the `rune` type.
- since runes are just an alias for int32, you're actually just working with numbers whenever you're working with a rune. This includes when you try to
  print out messages using runes. In order to print out the actual character it represents, you'll have to use a special formatting operation.

## rune byte representation: <<<<TILL HERE>>>>
Let's take a look at how runes are represented in memory:
![img.png](/img/section-2/7-006-4.png)
These 2 green rectangles represent int32 s. Each one of the blocks inside there, are 8 bits which for each green rectangle adds up to 32bits.
The number under each block is the number of it. So first one is 1, second is 2 and ... .

To represent capital letter Q, we have to have 51 in the first block(positon 1 byte).

To represent that symbol(zigzag), we have to have that combination of bytes in those specific locations in order to get a zigzag.

The reason this is important is because the letter Q is only 1 byte. The zigzag is 3 bytes. But if we were try to iterate over each letter, we'll actually
be iterating over each one of those bytes(like 51 in Q, or E2, 86 and AF in zigzag) instead of the entire letter itself(like the letter Q, so we don't iterate over Q
itself, but it's byte representation).

In future, we'll see how to iterate over the actual representations(the letter or characters themselves), instead of the byte values(like E2, 86, AF).

So bytes are **not** symbols as we seen.

![img.png](/img/section-2/7-006-5.png)

After runes, let's look at strings.

### string byte representation:
Let's look at how strings are represented in memory:
![img.png](/img/section-2/7-006-6.png)
Each one the runes has a specific bit length.

The string is a combination of all those runes on the left. That particular string is 0 byte long: 3 bytes from bitcoin sign, 1 byte for the 3, 3 bytes from
right arrow and 2 bytes from the cent sign. 

So if we were to iterate through the string, we'd be going 1 byte at a time and if you split those bytes in half, like when we're iterating over first byte(E2),
then you won't have the symbol meaning anymore because you're only iterating over 1 byte at a time.

Recall that runes are just numbers and since a string is composed of runes, a string is just more numbers put together.

To create runes in code, you surround the rune that you want to make with single quotes, so: 'a' is rune a.

We can create an omega symbol with backtick. So instead of using single quotes, we're using backticks on those characters in picture.
We use the backtick symbol whenever we're trying to type a symbol that isn't present on your keyboard. So basically for a-z,1-9 and 0 and any of
those symbols such as @, $, % and ..., it's ok to use single quotes, however everything else need to be in backticks.

A one exception is the escape symbols. The escape symbols start with a backslash. So \ + n -> will turn it into a new line.

To create strings, we use double quotes. 

We can also put a single letter within a string(double quotes) and that would be a string instead of a rune. Yeah, even though we could make it a rune with single
quote(because it's a single letter), since we want a string in the second line, we can use double quotes for that single letter.

You want to use raw literals whenever you're working with strings that have single quotes or double quotes in them.

The reason you want to use backticks for raw literals, is if you have double quotes within your string. Since the double quotes is used to create strings, if
you want to include them in your string, either have to escape them with the \ or instead of escaping those double quotes inside your string which also has double quotes,
you can just use a raw literal by using backticks. It's good to use a raw literal. Since it's easier to read.

Everything you include inside a raw literal(everything literal as you see it!), is gonna come out as is. That includes escape sequences like `\n`. So when you have
raw literal and you use a escape sequence inside of it, it's gonna come out as is and won't create a new line. So we're gonna get a backsla sh and n in the output instead of
a new line. 

So raw literals doesn't apply any processing to your string.

![img.png](/img/section-2/7-006-7.png)

![img.png](/img/section-2/7-006-8.png)

## 8-007 Go CLI:
![img.png](/img/section-2/8-007-1.png)
![img.png](/img/section-2/8-007-2.png)

fmt command is automated with your ide, but you can also run it manually.

## 9-008 Monthly Coding Challenges, Free Resources and Guides