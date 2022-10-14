//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string // type alias for library member name

type LendAudit struct {
	checkout time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total books owned by the library
	lended int // total books lended out
}

/* Library is gonna keep track of members  */
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string

		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkout.String(), "through", returnTime)
	}
}

// print all the audits for all the members
func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

/* A function that goes into the library prints how many books there are and how many have been lended out?  */
func printLibraryBooks(library *Library) {
	// start with a blank line, so it looks cleaner
	fmt.Println()

	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}

	fmt.Println()
}

/* The return type indicates that if the checkout was successful. */
func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	// here, we know we do have at least one book to lend
	book.lended += 1

	/* We're updating the book that we pulled out of library, so we need ot then reassign it back into the library, otherwise, we'll just
	be operating on the copy(the book variable that we pulled out at the beginning of the function). */
	library.books[title] = book
	member.books[title] = LendAudit{checkout: time.Now()} // `checkin` time will have a default value of 0

	return true // checkout was successful
}

// return a bool indicating this operation was successful or not
func returnBook(library *Library, title Title, member *Member) bool {
	/* First we need to make sure the book even exists in the library at all: */
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	// update the library:
	book.lended -= 1
	library.books[title] = book

	// update the checkIn info on the member:
	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The little Go book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susanna"] = Member{"Susanna", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheckout a book:")

	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
