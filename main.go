package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Заказ
	myOrder := Order{
		ID:      1,
		Item:    "Latte",
		Price:   300.0,
		IsReady: false,
	}

	myOrder.ApplyDiscount(15) // Скидка
	myOrder.FinishOrder()

	data, err := json.MarshalIndent(myOrder, "", "  ") 
	if err != nil {
		log.Fatal("Ошибка при создании JSON:", err)
	}

	// Результат
	fmt.Println("--- Итоговый чек (JSON) ---")
	fmt.Println(string(data))
}