package main

import (
	"errors"
	"fmt"
)

// Структура заказа
type Order struct {
	ID      int     `json:"order_id"`
	Item    string  `json:"item_name"`
	Price   float64 `json:"price"`
	IsReady bool    `json:"is_ready"`
}

func (o *Order) ApplyDiscount(percent float64) {
	o.Price -= o.Price * (percent / 100)
	fmt.Printf("📉 Скидка %.0f%% применена. Новая цена: %.2f\n", percent, o.Price)
}

func (o *Order) FinishOrder() {
	o.IsReady = true
}

type PaymentMethod interface {
	Pay(amount float64) error
}

// Оплата картой
type Card struct {
	Number   string
	CardDate string
	CardCVV  string
}

func (c Card) Pay(amount float64) error {
	if len(c.Number) != 16 {
		return errors.New("Неверный номер карты (нужно 16 цифр)")
	}

	if len(c.CardDate) != 4 {
		return errors.New("Неверный срок действия (нужно 4 цифры, например 1225)")
	}

	if len(c.CardCVV) != 3 {
		return errors.New("Неверный CVV (нужно 3 цифры)")
	}

	fmt.Printf("Карта %s: Списано %.2f руб. (Дата: %s, CVV: ***) OK\n", c.Number, amount, c.CardDate)
	return nil
}

// Оплата QR
type QRCode struct{ BankName string }

func (q QRCode) Pay(amount float64) error {
	fmt.Printf("QR %s: -%.2f руб. (OK)\n", q.BankName, amount)
	return nil
}