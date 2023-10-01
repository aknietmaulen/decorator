package main

import "fmt"

func main() {

	icecream := &MintIceCream{}

	//Add chocolate chips
	iceCreamWithChocolate := &ChocolateChips{
		icecream: icecream,
	}

	//Add strawberry topping
	iceCreamWithChocolateAndStrawberrySyrop := &StrawberrySyrop{
		icecream: iceCreamWithChocolate,
	}

	fmt.Printf("Price of mint ice cream with chocolate and strawberry topping is %d\n", iceCreamWithChocolateAndStrawberrySyrop.getPrice())
}
