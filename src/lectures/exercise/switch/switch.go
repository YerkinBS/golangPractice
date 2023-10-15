//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {
	var age int
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении ввода:", err)
        return
	}
	
	switch {
	case age == 0:
		fmt.Println("newborn")
	case 1 <= age && age <= 3:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
