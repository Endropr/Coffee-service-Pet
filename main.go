package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const ShopName = "Coffee Shop"

func main() {
	fmt.Printf("=== %s ===\n\n", ShopName)

	myOrder := Order{ID: 1, Item: "Latte", Price: 300.0}
	myOrder.ApplyDiscount(15)
	myOrder.FinishOrder()

	// Ввод номера карты
	var cardNumber string
	var cardDate string
	var cardCVV string
	for {
		fmt.Print("Введите 16 цифр номера карты: ")
		fmt.Scan(&cardNumber)

		fmt.Print("Введите срок действия карты (4 цифры): ")
		fmt.Scan(&cardDate)

		fmt.Print("Введите CVV карты (3 цифры):")
		fmt.Scan(&cardCVV)

		payment := Card{
			Number:   cardNumber,
			CardDate: cardDate,
			CardCVV:  cardCVV,
		}

		err := executePayment(myOrder.Price, payment)
		if err == nil {
			break
		}

		fmt.Printf("Попробуйте еще раз...\n")
	}

	saveReceipt(myOrder)
}

func executePayment(amount float64, method PaymentMethod) error {
	if err := method.Pay(amount); err != nil {
		fmt.Println("Ошибка:", err)
		return err
	}
	fmt.Println("Оплата прошла успешно!")
	return nil
}

// Запись в json
func saveReceipt(order Order) {
	data, _ := json.MarshalIndent(order, "", "  ")

	fmt.Printf("\n--- Итоговый чек ---\n%s\n", string(data))

	if err := os.WriteFile("order.json", data, 0644); err != nil {
		log.Printf("Внимание! Ошибка записи: %v", err)
	} else {
		fmt.Println("\n💾 Чек сохранен в order.json")
	}
}
