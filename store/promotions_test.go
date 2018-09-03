package store

import (
	"testing"
)

func TestCart_XForPriceOfY(t *testing.T) {
	type fields struct {
		Items map[string]int
	}
	type args struct {
		x  int
		y  int
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			"Zero Google Homes at three for the price of two",
			fields{
				Items: map[string]int{
					"120P90": 0,
				},
			},
			args{3, 2, "120P90"},
			0,
		},
		{
			"Three Google Homes at three for the price of two",
			fields{
				Items: map[string]int{
					"120P90": 3,
				},
			},
			args{3, 2, "120P90"},
			49.99,
		},
		{
			"Two Google Homes at three for the price of two",
			fields{
				Items: map[string]int{
					"120P90": 2,
				},
			},
			args{3, 2, "120P90"},
			0,
		},
		{
			"Four Google Homes at three for the price of two",
			fields{
				Items: map[string]int{
					"120P90": 4,
				},
			},
			args{3, 2, "120P90"},
			49.99,
		},
		{
			"Six Google Homes at three for the price of two",
			fields{
				Items: map[string]int{
					"120P90": 6,
				},
			},
			args{3, 2, "120P90"},
			99.98,
		},
	}
	for _, tt := range tests {
		SetStock()

		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				Items: tt.fields.Items,
			}
			if got := c.XForPriceOfY(tt.args.x, tt.args.y, tt.args.id); got != tt.want {
				t.Errorf("Cart.XForPriceOfY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_OverXDiscountY(t *testing.T) {
	type fields struct {
		Items map[string]int
	}
	type args struct {
		x  int
		y  float32
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			"Zero Alexa Speakers at 10% off over 3",
			fields{
				Items: map[string]int{
					"A304SD": 0,
				},
			},
			args{3, 10, "A304SD"},
			0,
		},
		{
			"Three Alexa Speakers at 10% off over 3",
			fields{
				Items: map[string]int{
					"A304SD": 3,
				},
			},
			args{3, 10, "A304SD"},
			0,
		},
		{
			"Four Alexa Speakers at 10% off over 3",
			fields{
				Items: map[string]int{
					"A304SD": 4,
				},
			},
			args{3, 10, "A304SD"},
			43.8,
		},
	}
	for _, tt := range tests {
		SetStock()

		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				Items: tt.fields.Items,
			}
			if got := c.OverXDiscountY(tt.args.x, tt.args.y, tt.args.id); got != tt.want {
				t.Errorf("Cart.OverXDiscountY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_BuyXGetYFree(t *testing.T) {
	type fields struct {
		Items map[string]int
	}
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
		want1  string
	}{
		{
			"One MacBook Pro, Zero Rasberry Pi B",
			fields{
				Items: map[string]int{
					"43N23P": 1,
				},
			},
			args{"43N23P", "234234"},
			0,
			"You have recieved 1 free Rasberry Pi B(s) along with your purchase of MacBook Pro",
		},
		{
			"One MacBook Pro, One Rasberry Pi B",
			fields{
				Items: map[string]int{
					"43N23P": 1,
					"234234": 1,
				},
			},
			args{"43N23P", "234234"},
			30,
			"You have recieved 1 free Rasberry Pi B(s) along with your purchase of MacBook Pro",
		},
		{
			"One MacBook Pro, Two Rasberry Pi B",
			fields{
				Items: map[string]int{
					"43N23P": 1,
					"234234": 2,
				},
			},
			args{"43N23P", "234234"},
			30,
			"You have recieved 1 free Rasberry Pi B(s) along with your purchase of MacBook Pro",
		},
		{
			"Two MacBook Pro, Three Rasberry Pi B",
			fields{
				Items: map[string]int{
					"43N23P": 2,
					"234234": 3,
				},
			},
			args{"43N23P", "234234"},
			60,
			"You have recieved 2 free Rasberry Pi B(s) along with your purchase of MacBook Pro",
		},
		{
			"Five MacBook Pro, Two Rasberry Pi B (Not enough stock)",
			fields{
				Items: map[string]int{
					"43N23P": 5,
					"234234": 2,
				},
			},
			args{"43N23P", "234234"},
			60,
			"You have recieved 2 free Rasberry Pi B(s) along with your purchase of MacBook Pro\nUnfortunatly we are out of stock to add the remaining free items",
		},
	}
	for _, tt := range tests {
		SetStock()

		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				Items: tt.fields.Items,
			}
			got, got1 := c.BuyXGetYFree(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("Cart.BuyXGetYFree() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Cart.BuyXGetYFree() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
