package main

import (
	"fmt"
)

// das Konto — банківський рахунок
type Konto struct {
	Guthaben float64 // das Guthaben — баланс
}

// die Einzahlung — поповнення рахунку
// der Betrag — сума
func (k *Konto) GeldEinzahlen(betrag float64) error {
	if betrag <= 0 {
		// der Fehler — помилка
		return fmt.Errorf("der Betrag muss groesser als 0 sein")
	}
	k.Guthaben += betrag
	return nil
}

// die Abhebung — зняття коштів
func (k *Konto) GeldAbheben(betrag float64) error {
	if betrag <= 0 {
		return fmt.Errorf("der Betrag muss groesser als 0 sein")
	}
	if betrag > k.Guthaben {
		// nicht genug Geld — недостатньо коштів
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %.2f abzuziehen", betrag)
	}
	k.Guthaben -= betrag
	return nil
}

func main() {
	// das Konto — банківський рахунок
	meinKonto := Konto{
		Guthaben: 100.0, // das Guthaben — баланс
	}

	// die Einzahlung — поповнення
	if err := meinKonto.GeldEinzahlen(50.0); err != nil {
		// der Fehler — помилка
		fmt.Println("Fehler bei der Einzahlung:", err)
	}
	fmt.Printf("Aktuelles Guthaben: %.2f EUR\n", meinKonto.Guthaben) // aktuell — поточний

	// die Abhebung — зняття
	err := meinKonto.GeldAbheben(200.0)
	if err != nil {
		fmt.Println(err)
	}
}
