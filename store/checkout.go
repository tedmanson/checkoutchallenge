package store

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Checkout prints the final listing of contents and the total price of the shopping cart
func (c Cart) Checkout() {
	var total = c.ScannedItemsTotalPrices()
	discount, messages := c.ApplyPromotions()
	total -= discount

	p := message.NewPrinter(language.English)
	p.Printf("Scanned items: %s\n", c.ScannedItemsLabels())
	if len(messages) > 0 {
		p.Printf("Note: %s\n", messages)
	}
	p.Printf("Total: $%.2f\n", total)

}
