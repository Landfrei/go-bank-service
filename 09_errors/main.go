package main

import (
	"09_errors/bank"
	"fmt"
)

func main() {
	konto := &bank.Konto{Guthaben: 500}
	err := konto.GeldAbheben(-50)
	if err != nil {
		fmt.Println("Fehler:", err) // der Fehler - помилка
	} else {
		fmt.Println("Neues Guthaben:", konto.Guthaben)
	}
	err = konto.GeldAbheben(600)
	if err != nil {
		fmt.Println("Fehler:", err) // der Fehler - помилка
	} else {
		fmt.Println("Neues Guthaben:", konto.Guthaben)
	}
	err = konto.GeldAbheben(200)
	if err == nil {
		fmt.Println("Erfolg! Neuer Kontostand:", konto.Guthaben) // der Erfolg - успіх, der Kontostand - баланс рахунку
	}
}
