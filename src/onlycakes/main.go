package main

import (
	"fmt"
	models "onlycakes/models"
)

var cakes = []models.Option{
	{Id: 1, Name: "Basic Cake", Price: 10.00},
	{Id: 2, Name: "Two-Layer Cake", Price: 12.00},
	{Id: 3, Name: "Three-Layer Cake", Price: 15.00},
	{Id: 4, Name: "Cupcake", Price: 3.75}}

var cakeFlavorOptions = []models.Option{
	{Id: 1, Name: "Chocolate", Price: .00},
	{Id: 2, Name: "Vanilla", Price: 0.00},
	{Id: 3, Name: "Spice", Price: 1.00},
	{Id: 4, Name: "Coffee", Price: 2.00},
}

var toppingOptions = []models.Option{
	{Id: 1, Name: "Vanilla", Price: 1.00},
	{Id: 2, Name: "Chocolate", Price: 2.00},
	{Id: 3, Name: "Strawberry", Price: 3.00},
	{Id: 4, Name: "Ganache", Price: 5.00},
	{Id: 0, Name: "None", Price: 0.00},
}

var order = models.Order{}

func main() {
	fmt.Println("\nWelcome to OnlyCakes!")

	var optionSelected uint8

	for {
		printOptionMenu()
		fmt.Print("\n>: ")
		fmt.Scan(&optionSelected)

		if optionSelected == 0 {
			fmt.Println("\nThanks come again!")
			break
		}

		handleEnteredOption(optionSelected)
	}
}

func handleEnteredOption(option uint8) {
	switch option {
	case 1:
		printCustomMenu(cakes)
	case 2:
		handleOrder()
	case 3:
		printOrder()
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
	var selectedToppingFlavorId uint8

	fmt.Print("\nCake Option: ")
	fmt.Scan(&selectedCakeOptionId)

	// Get cake flavor
	printCustomMenu(cakeFlavorOptions)
	fmt.Print("\nCake Flavor: ")
	fmt.Scan(&selectedCakeFlavorId)

	// Get topping flavor
	printCustomMenu(toppingOptions)
	fmt.Print("\nTopping Flavor: ")
	fmt.Scan(&selectedToppingFlavorId)

	// Validate user input
	isValid, errorMessage := validateUserInput(selectedCakeOptionId, selectedCakeFlavorId, selectedToppingFlavorId)
	if !isValid {
		fmt.Printf("\n%s\n", errorMessage)
		return
	}

	fmt.Println("\nYou selected the following...")

	// Calculate cost
	cake := cakes[selectedCakeOptionId-1]
	cakeFlavor := cakeFlavorOptions[selectedCakeFlavorId-1]

	var toppingFlavor models.Option
	if selectedToppingFlavorId != 0 {
		toppingFlavor = toppingOptions[selectedToppingFlavorId-1]
	} else {
		// Set to last option ("None")
		toppingFlavor = toppingOptions[len(toppingOptions)-1]
	}

	total := cake.Price + cakeFlavor.Price + toppingFlavor.Price
	newItem := models.Item{CakeType: cake, CakeFlavor: cakeFlavor, ToppingFlavor: toppingFlavor, Total: total}

	fmt.Printf("\n%s\n", getItemDescription(newItem))

	order.AddItem(newItem)

	fmt.Printf("\nItem Total: $%.2f\n", newItem.Total)
	fmt.Printf("\nTotal Items: %d\n", len(order.Items))
	fmt.Printf("Order Total: $%.2f\n", order.Total)
}

func getItemDescription(item models.Item) string {
	return fmt.Sprintf("%s %s with %s frosting", item.CakeFlavor.Name, item.CakeType.Name, item.ToppingFlavor.Name)
}

func validateUserInput(cakeOption uint8, cakeFlavor uint8, toppingFlavor uint8) (bool, string) {
	if cakeOption-1 < 0 || (int)(cakeOption-1) > len(cakes)-1 {
		return false, "Invalid cake option."
	}

	if cakeFlavor-1 < 0 || (int)(cakeFlavor-1) > len(cakeFlavorOptions)-1 {
		return false, "Invalid cake flavor option."
	}

	if toppingFlavor-1 < 0 || (int)(toppingFlavor-1) > len(toppingOptions)-1 {
		return false, "Invalid topping option."
	}

	return true, ""
}

func printCustomMenu(options []models.Option) {
	fmt.Println("\nOur current offerings:")
	for _, cake := range options {
		fmt.Printf("%d. %s $%.2f\n", cake.Id, cake.Name, cake.Price)
	}
}

func printOrder() {
	fmt.Println("\n========================================================++")
	fmt.Println("\nYour current order:")
	for i, item := range order.Items {
		fmt.Printf("\nItem %d: %s $%.2f", i+1, getItemDescription(item), item.Total)
	}

	fmt.Printf("\n\nOrder Total: %.2f\n", order.Total)
	fmt.Println("\n========================================================++")
}

func printOptionMenu() {
	fmt.Println("\nEnter an option from the menu to get started.")
	fmt.Println("1. View the menu")

	if len(order.Items) == 0 {
		fmt.Println("2. Start an order")
	} else {
		fmt.Println("2. Add Item")
		fmt.Println("3. View Order")
	}

	fmt.Println("0. Quit")
}
