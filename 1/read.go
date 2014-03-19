package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the pace calculator\n");
	fmt.Println("Please enter distances in miles and times in the format"
	fmt.Println("hours:minutes:seconds, as in 0:32:27.\n")

	fmt.Printf("Distance? ")
	var distance float64
	fmt.Scanf("%f", &distance)
	fmt.Println("distance is", distance)
}
