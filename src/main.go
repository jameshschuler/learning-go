package main

import "fmt"

type Cake struct {
	name  string
	price float32
}

var cakes = []Cake{
	{name: "Basic Cake", price: 10.00},
	{name: "Two-Layer Cake", price: 12.00},
	{name: "Three-Layer Cake", price: 15.00},
	{name: "Cupcake", price: 3.75}}

func main() {
	fmt.Println("\nWelcome to OnlyCakes!")
	printOptionMenu()

	var optionSelected uint8

	for {
		fmt.Print("\nMenu Option: ")
		fmt.Scan(&optionSelected)

		if optionSelected == 0 {
			fmt.Println("Thanks come again!")
			break
		}

		handleEnteredOption(optionSelected)
	}
}

func handleEnteredOption(option uint8) {
	switch option {
	case 1:
		handleOrder()
	case 2:
		printCakesMenu()
	default:
		fmt.Printf("Option %v not supported.\n", option)
		break
	}
}

func handleOrder() {
	fmt.Print("\nEnter which cake you would like.")
	printCakesMenu()

	// Get cake option
	var selectedCakeOption uint8

	fmt.Print("\nCake Option: ")
	fmt.Scan(&selectedCakeOption)

	// Get cake flavor

	// Get frosting flavor

	// Calculate cost
}

func printCakesMenu() {
	fmt.Println("\nOur current offerings:")
	for i, cake := range cakes {
		fmt.Printf("%d. %s $%.2f\n", i+1, cake.name, cake.price)
	}
}

func printOptionMenu() {
	fmt.Println("Enter an option from the menu to get started.")
	fmt.Println("1. Place an order")
	fmt.Println("2. View the menu")
	fmt.Println("0. Quit")
}
