package main

import (
	"07_pointers_receivers/bank"
	"fmt"
)

func main() {
	// das Muster — зразок / шаблон
	// Max Mustermann — Макс Шаблоненко
	// das Konto — банківський рахунок
	k := bank.Konto{Inhaber: "Max Mustermann", Guthaben: 100.0} // der Inhaber — власник, das Guthaben — баланс
	fmt.Printf("Startguthaben: %.2f EUR\n", k.Guthaben)

	// die Einzahlung nach Wert — поповнення через копію
	// der Wert — копія
	k.EinzahlenWert(50.0)
	fmt.Printf("Guthaben nach EinzahlenWert: %.2f EUR\n", k.Guthaben) // nach — після, nach EinzahlenWert — після поповнення через копію

	// der Zeiger — вказівник
	k.EinzahlenZeiger(50.0)
	fmt.Printf("Guthaben nach EinzahlenZeiger: %.2f EUR\n", k.Guthaben) // nach — після, nach EinzahlenZeiger — після поповнення через вказівник
}
