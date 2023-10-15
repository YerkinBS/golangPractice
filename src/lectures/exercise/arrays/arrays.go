//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price int
}

func main() {
	shoppingList := [...]Product {
		{name: "Skirt", price: 50},
		{name: "Headphones", price: 300},
		{name: "PS5", price: 900},
	}

	fmt.Println("The last item on the list:", shoppingList[len(shoppingList) - 1])
	fmt.Println("The total number of items:", len(shoppingList))

	var cnt int
	for i := 0; i < len(shoppingList); i++ {
		cnt += shoppingList[i].price
	}
	fmt.Println("The total cost of the items:", cnt)
}
