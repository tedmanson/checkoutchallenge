package store

// NoItemsInCart is the fixed message for logging when a cart has zero items
const NoItemsInCart = "No items in shopping cart"

// OutOfStock is the fixed message for logging when an item is out of stock
const OutOfStock = "Out of stock"

// ZeroItemCountInCart is the fixed message for logging when a cart has an item with zero items added
const ZeroItemCountInCart = "Item with zero count in cart"

func init() {
	inv.ProductListing["120P90"] = Product{
		Details: ProductDetails{
			SKU:         "120P90",
			ProductName: "Google Home",
			ListedPrice: 49.99,
		},
		Stock: 10,
	}

	inv.ProductListing["43N23P"] = Product{
		Details: ProductDetails{
			SKU:         "43N23P",
			ProductName: "MacBook Pro",
			ListedPrice: 5399.99,
		},
		Stock: 5,
	}

	inv.ProductListing["A304SD"] = Product{
		Details: ProductDetails{
			SKU:         "A304SD",
			ProductName: "Alexa Speaker",
			ListedPrice: 109.50,
		},
		Stock: 10,
	}

	inv.ProductListing["234234"] = Product{
		Details: ProductDetails{
			SKU:         "234234",
			ProductName: "Rasberry Pi B",
			ListedPrice: 30.00,
		},
		Stock: 2,
	}

}
