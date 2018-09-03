package store

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Cart contains a map of items to item count
type Cart struct {
	Items map[string]int
}

// New returns a fully initilized shopping cart
func New() *Cart {
	return &Cart{
		Items: make(map[string]int),
	}
}

// AddItem adds a product id to the shopping cart. If one already exists it increments this number
func (c Cart) AddItem(id string) error {
	if _, ok := inv.ProductListing[id]; !ok {
		return fmt.Errorf("Product does not exist in inventory: %s", id)
	}

	p := inv.ProductListing[id]
	if p.Stock < 1 {
		log.WithFields(log.Fields{
			"cause": "ScannedItemsLabels",
			"sku":   id,
		}).Info(OutOfStock)

		return fmt.Errorf("Not enough stock of %s (%s)", p.Details.ProductName, p.Details.SKU)
	}

	p.Stock--
	inv.ProductListing[id] = p

	var count = 1
	if _, ok := c.Items[id]; ok {
		count = c.Items[id] + 1
	}

	c.Items[id] = count

	return nil
}

// ScannedItemsLabels returns a slice of product names from the shopping cart
func (c Cart) ScannedItemsLabels() string {
	if len(c.Items) == 0 {
		return NoItemsInCart
	}

	var items []string
	for id, count := range c.Items {
		if count == 0 {
			log.WithFields(log.Fields{
				"cause": "ScannedItemsLabels",
				"sku":   id,
			}).Info(ZeroItemCountInCart)
			continue
		}

		if p, ok := inv.ProductListing[id]; ok {
			if count > 1 {
				items = append(items, p.Details.ProductName+" x"+strconv.Itoa(count))
			} else {
				items = append(items, p.Details.ProductName)
			}

		} else {
			log.WithFields(log.Fields{
				"cause": "ScannedItemsLabels",
				"sku":   id,
			}).Info("Item missing from inventory")

		}
	}

	if len(items) == 0 {
		return NoItemsInCart
	}

	sort.Strings(items)
	return strings.Join(items, ", ")
}

// ScannedItemsTotalPrices returns the float32 sum of all products in the shopping cart
func (c Cart) ScannedItemsTotalPrices() float32 {
	var total float32

	for id, count := range c.Items {
		if d, ok := inv.ProductListing[id]; ok {
			total = total + (d.Details.ListedPrice * float32(count))
		} else {
			log.WithFields(log.Fields{
				"cause": "ScannedItemsLabels",
				"sku":   id,
			}).Info("Item missing from inventory")
		}
	}

	return total
}
