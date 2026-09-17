package main

import (
	"09_errors/bank"
	"errors"
	"fmt"
)

func main() {
	konto := &bank.Konto{Guthaben: 500}
	err := konto.GeldAbheben(600)
	if err != nil {
		var fundsErr *bank.UnzureichendesGuthabenFehler
		if errors.As(err, &fundsErr) {
			fmt.Println("=== Спрацювала німецька кастомна помилка! ===")
			fmt.Printf("Запитано (Betrag): %.2f EUR\n", fundsErr.Betrag)
			fmt.Printf("Доступно (Guthaben): %.2f EUR\n", fundsErr.Guthaben)
			fmt.Printf("Дефіцит: %.2f EUR\n", fundsErr.Betrag-fundsErr.Guthaben)
		} else {
			fmt.Println("=== Спрацювала стандартна помилка! ===")
			fmt.Println(err)
		}
	}
}
