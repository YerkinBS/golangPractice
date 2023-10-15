//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length, width int
}

func Area(length, width int) int {
	return length * width
}

func Perimeter(length, width int) int {
	return 2 * (length + width)
}


func main() {
	rect := Rectangle {5, 10}

	fmt.Println("Area of the rectangle:", Area(rect.length, rect.width))
	fmt.Println("Perimeter of the rectangle:", Perimeter(rect.length, rect.width))

	rect.length *= 2
	rect.width *= 2
	fmt.Println("Area of the doubled rectangle:", Area(rect.length, rect.width))
	fmt.Println("Perimeter of the doubled rectangle:", Perimeter(rect.length, rect.width))
}
