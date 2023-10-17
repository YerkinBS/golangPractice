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
	"math/rand"
	"time"
)

type Status struct {
	checkedOut string
	checkedIn  string
	member     string
}

type Book struct {
	name   string
	status Status
}

type Library struct {
	members []string
	books   []Book
}

func getPos(lib *Library, bookName string) (bool, int) {
	bookFound, bookPos := false, -1
	for i, book := range lib.books {
		if book.name == bookName {
			bookFound = true
			bookPos = i
		}
	}
	return bookFound, bookPos
}

func timeFormatChanger(timeNow time.Time) string {
	return timeNow.Format("2006-01-02 15:04:05")
}

func checkOut(lib *Library, bookName string) {
	bookFound, bookPos := getPos(lib, bookName)

	if !bookFound {
		fmt.Println("\nThis book is not registered in the library. Try to choose another book. Thanks!")
	} else {
		if lib.books[bookPos].status.checkedOut == "" {
			lib.books[bookPos].status.checkedOut = timeFormatChanger(time.Now())
			lib.books[bookPos].status.checkedIn = ""
			lib.books[bookPos].status.member = lib.members[rand.Intn(len(lib.members))]
			fmt.Println("\nThe book named", bookName, "is checked out succesfully by", lib.books[bookPos].status.member)
		} else {
			fmt.Println("\nThe book named", bookName, "has already been checked out by", lib.books[bookPos].status.member, "in", lib.books[bookPos].status.checkedOut)
		}
	}
}

func checkIn(lib *Library, bookName string) {
	bookFound, bookPos := getPos(lib, bookName)

	if !bookFound {
		fmt.Println("\nThis book is not registered in the library. Try to choose another book. Thanks!")
	} else {
		if lib.books[bookPos].status.checkedIn == "" {
			lib.books[bookPos].status.checkedIn = timeFormatChanger(time.Now())
			lib.books[bookPos].status.checkedOut = ""
			lib.books[bookPos].status.member = lib.members[rand.Intn(len(lib.members))]
			fmt.Println("\nThe book named", bookName, "is checked in succesfully by", lib.books[bookPos].status.member)
		} else {
			fmt.Println("\nThe book named", bookName, "has already been checked in by", lib.books[bookPos].status.member, "in", lib.books[bookPos].status.checkedIn)
		}
	}
}

func main() {
	library := Library{
		members: []string{"Yerkin", "Aidana", "Nurtay", "Ambal"},
		books: []Book{
			{name: "Eternal Echoes: A Time Travel Adventure"},
			{name: "The Enchanted Forest Chronicles"},
			{name: "The Midnight Library"},
			{name: "The Lost Symbol"},
		},
	}

	checkOut(&library, "The Lost Symbol")
	fmt.Println("\n", library)
	checkIn(&library, "The Midnight Library")
	fmt.Println("\n", library)
	checkIn(&library, "The Midnight Library")
	checkOut(&library, "The Lost Symbol")
}
