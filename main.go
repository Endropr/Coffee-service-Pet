package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	
	err = os.WriteFile("order.json", data, 0644)
	if err != nil {
		log.Printf("Внимание! Не удалось сохранить файл: %v", err)
	} else {
		fmt.Println("\n✅ Чек успешно сохранен в файл: order.json")
	}
}