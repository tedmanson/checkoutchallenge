package store

var inv = Inventory{
	ProductListing: make(map[string]Product),
}

// Inventory is a map of sku to product
type Inventory struct {
	ProductListing map[string]Product
}

// Product is a container for a product item and its stock level
type Product struct {
	Details ProductDetails
	Stock   int
}

// ProductDetails contains the product details
type ProductDetails struct {
	SKU         string
	ProductName string
	ListedPrice float32
}
