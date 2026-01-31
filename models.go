package main

import "fmt"

// Структура
type Order struct {
	ID      int     `json:"id"`
	Item    string  `json:"item_name"`
	Price   float64 `json:"price"`
	IsReady bool    `json:"is_ready"`
}
func (o *Order) ApplyDiscount(percent float64) {
	o.Price = o.Price * (1 - percent/100)
	fmt.Printf("Скидка применилась! Новая цена: %.2f\n", o.Price)
}

func (o *Order) FinishOrder() {
	o.IsReady = true
}
