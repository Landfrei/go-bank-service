package main

import (
	"fmt"
)

// das Konto — обліковий запис / банківський рахунок
type Konto struct {
	Guthaben float64 // das Guthaben — баланс / кошти
}

// die Einzahlung — поповнення рахунку
func (k *Konto) GeldEinzahlen(betrag float64) error {
	// der Betrag — сума
	if betrag <= 0 {
		return fmt.Errorf("der Betrag muss größer als 0 sein")
	}
	k.Guthaben += betrag
	return nil
}

// die Abhebung — зняття коштів
func (k *Konto) GeldAbheben(betrag float64) error {
	if betrag <= 0 {
		return fmt.Errorf("der Betrag muss größer als 0 sein")
	}
	if betrag > k.Guthaben {
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %.2f Geld abzuziehen", betrag)
	}
	k.Guthaben -= betrag
	return nil
}

func main() {
	meinBank := Konto{
		Guthaben: 100.0,
	}

	// Поповнення
	if err := meinBank.GeldEinzahlen(50.0); err != nil {
		fmt.Println("Fehler bei der Einzahlung:", err)
	}
	fmt.Printf("Mein Guthaben ist %.2f Geld\n", meinBank.Guthaben)

	// Спроба зняти більше, ніж є на рахунку
	err := meinBank.GeldAbheben(200.0)
	if err != nil {
		fmt.Println(err)
	}
}
