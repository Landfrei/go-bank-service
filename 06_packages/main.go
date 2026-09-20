package main

import (
	"06_packages/bank"
	"fmt"
)

func main() {
	// Max Mustermann — Макс Шаблоненко, das Muster — зразок / шаблон
	konto := bank.Konto{Inhaber: "Max Mustermann", Guthaben: 100} // das Konto — банківський рахунок

	err := bank.Abwickeln(&konto, 40) // die Abwicklung — проведення платежу
	if err != nil {
		fmt.Println("Fehler: ", err) // der Fehler — помилка
	}
	fmt.Println("Guthaben: ", konto.Guthaben) // das Guthaben — баланс

	err = bank.Abwickeln(&konto, 7676)
	if err != nil {
		fmt.Println("Fehler: ", err)
	}
	fmt.Println("Guthaben: ", konto.Guthaben)
}
