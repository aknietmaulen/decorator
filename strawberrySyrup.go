package main

type StrawberrySyrop struct {
	icecream IceCream
}

func (c *StrawberrySyrop) getPrice() int {
	iceCreamPrice := c.icecream.getPrice()
	return iceCreamPrice + 50
}
