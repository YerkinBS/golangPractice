//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	active = true
	inactive = false
)

type Item struct {
	name string
	securityTag bool
}

func activate(item *Item) {
	item.securityTag = true
}

func deactivate(item *Item) {
	item.securityTag = false
}

func checkout(itemsList []Item) {
	for i := 0; i < len(itemsList); i++ {
		deactivate(&itemsList[i])
	}
}

func main() {
	stool, rope, soap, paper := Item {"Stool", active}, Item {"Rope", active}, Item {"Soap", active}, Item {"Paper", active}
	itemSlice := []Item{stool, rope, soap, paper}

	fmt.Println("Initial", itemSlice)
	deactivate(&itemSlice[2])
	fmt.Println("Item 3 deactivated", itemSlice)
	checkout(itemSlice)
	fmt.Println("Checked out", itemSlice)
	activate(&itemSlice[3])
	fmt.Println("Item 4 activated", itemSlice)
}
