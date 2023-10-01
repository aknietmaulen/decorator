package main

type ChocolateChips struct {
	icecream IceCream
}

func (c *ChocolateChips) getPrice() int {
	iceCreamPrice := c.icecream.getPrice()
	return iceCreamPrice + 70
}
