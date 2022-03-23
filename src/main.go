package main

import "fmt"

type Option struct {
	id    uint8
	name  string
	price float32
}

var cakes = []Option{
	{id: 1, name: "Basic Cake", price: 10.00},
	{id: 2, name: "Two-Layer Cake", price: 12.00},
	{id: 3, name: "Three-Layer Cake", price: 15.00},
	{id: 4, name: "Cupcake", price: 3.75}}

var cakeFlavorOptions = []Option{
	{id: 1, name: "Chocolate", price: 0.00},
	{id: 2, name: "Vanilla", price: 0.00},
	{id: 3, name: "Spice", price: 1.00},
	{id: 4, name: "Coffee", price: 2.00},
}

var toppingOptions = []Option{
	{id: 1, name: "None", price: 0.00},
	{id: 2, name: "Vanilla", price: 0.00},
	{id: 3, name: "Chocolate", price: 1.00},
	{id: 4, name: "Strawberry", price: 2.00},
	{id: 5, name: "Ganache", price: 5.00},
}

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
		printCustomMenu(cakes)
	default:
		fmt.Printf("Option %v is not supported.\n", option)
		break
	}
}

func handleOrder() {
	printCustomMenu(cakes)

	// Get cake option
	var selectedCakeOptionId uint8
	var selectedCakeFlavorId uint8

	fmt.Print("\nCake Option: ")
	fmt.Scan(&selectedCakeOptionId)

	// Get cake flavor
	printCustomMenu(cakeFlavorOptions)
	fmt.Print("\nCake Flavor: ")
	fmt.Scan(&selectedCakeFlavorId)

	// Get frosting flavor

	// Calculate cost
}

func printCustomMenu(options []Option) {
	fmt.Println("\nOur current offerings:")
	for _, cake := range options {
		fmt.Printf("%d. %s $%.2f\n", cake.id, cake.name, cake.price)
	}
}

func printOptionMenu() {
	fmt.Println("Enter an option from the menu to get started.")
	fmt.Println("1. Place an order")
	fmt.Println("2. View the menu")
	fmt.Println("0. Quit")
}
