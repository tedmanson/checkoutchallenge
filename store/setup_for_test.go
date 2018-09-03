package store

func SetStock() {
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
