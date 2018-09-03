# Shopping Cart

Have you shopped online? Let’s imagine that you need to build the checkout backend service that will support different promotions with the given inventory.

## Initial Inventory

Build a checkout system with these items:

| SKU           | Name           | Price     | Inventory Qty |
| ------------- |:--------------:| ---------:|--------------:|
| 120P90        | Google Home    | $49.99    | 10            |
| 43N23P        | MacBook Pro    | $5,399.99 | 5             |
| A304SD        | Alexa Speaker  | $109.50   | 10            |
| 234234        | Raspberry Pi B | $30.00    | 2             |

## Intitial Promotions

The system should have the following promotions:

- Each sale of a MacBook Pro comes with a free Raspberry Pi B
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

## Example Scenarios

- Scanned Items: MacBook Pro, Raspberry Pi B Total: $5,399.99
- Scanned Items: Google Home, Google Home, Google Home Total: $99.98
- Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker Total: $295.65

## Usage

Current implementation is hardcoded to a Macbook Pro and a Rasberry Pi B. Updating the ID's in the AddItem function will change the result printed to the terminal.

```golang
var c = store.New()

c.AddItem(`234234`)
c.AddItem(`43N23P`)

c.Checkout()
```

## Improvements post build

This system is designed solely as a shopping cart. Everythign built within this should be pulled back out of this and into separate packages.

### Store

It would be ideal to in this case have user entered items. Be able to list the items available and their current stock levels. Currently it is a proof of concept and hardcoded.

### Inventory

There should be a separate service for retrieving from the inventory. Even in the current implemntation there are risks surrounding the "map" implementation as no sync package has been added to it to manage high concurrency. Maps are not threadsafe within Golang and need to Lock and Unlock if they are to maintain proper state.

### Promotions

Promotions would be created separately and as a separate service that is managed through an admin tool (programatic promotions such as is currently entered, but not hardcoding the promotions within).

### Responses

All responses should be correctly listed and maintained in consts or if user facing via an interface.