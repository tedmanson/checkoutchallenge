package main

import "github.com/tedmanson/checkoutchallenge/store"

func main() {
	var c = store.New()

	c.AddItem(`234234`)
	c.AddItem(`43N23P`)
	c.Checkout()
}
