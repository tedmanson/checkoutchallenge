package store

import (
	"testing"
)

// Tests within this do not incorporate any promotions. It is validating the correct summation of original prices.
func TestCart_ScannedItemsTotalPrices(t *testing.T) {

	type fields struct {
		Items map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{"No items in cart", fields{Items: map[string]int{}}, 0},

		{"Incorrect item in cart", fields{
			Items: map[string]int{
				"incorrect": 1,
			},
		}, 0},

		{"MacBook Pro and Rasberry Pi B item in cart", fields{
			Items: map[string]int{
				"43N23P": 1,
				"234234": 1,
			},
		}, 5429.99},

		{"Three Rasberry Pi B's item in cart", fields{
			Items: map[string]int{
				"234234": 3,
			},
		}, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				Items: tt.fields.Items,
			}
			if got := c.ScannedItemsTotalPrices(); got != tt.want {
				t.Errorf("Cart.ScannedItemsTotalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_ScannedItemsLabels(t *testing.T) {
	type fields struct {
		Items map[string]int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"No items in cart", fields{Items: map[string]int{}}, NoItemsInCart},

		{"Incorrect item in cart", fields{
			Items: map[string]int{
				"incorrect": 1,
			},
		}, NoItemsInCart},

		{"Rasberry Pi B in cart", fields{
			Items: map[string]int{
				"234234": 1,
			},
		}, `Rasberry Pi B`},

		{"Two Rasberry Pi B's in cart", fields{
			Items: map[string]int{
				"234234": 2,
			},
		}, `Rasberry Pi B x2`},

		{"Rasberry Pi B's removed from cart but not removed ID", fields{
			Items: map[string]int{
				"234234": 0,
			},
		}, NoItemsInCart},

		{"Rasberry Pi B and Google Home in cart", fields{
			Items: map[string]int{
				"234234": 1,
				"120P90": 1,
			},
		}, `Google Home, Rasberry Pi B`},
	}
	for _, tt := range tests {
		SetStock()

		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				Items: tt.fields.Items,
			}
			if got := c.ScannedItemsLabels(); got != tt.want {
				t.Errorf("Cart.ScannedItemsLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
