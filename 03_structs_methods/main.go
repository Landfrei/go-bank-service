package main

import (
	"fmt"
)

// das Konto — банківський рахунок
// das Guthaben — баланс
type Konto struct {
	IBAN     string
	Guthaben int
}

// der Kunde — клієнт
type Kunde struct {
	ID    int
	Name  string
	Konto *Konto
}

// die Prüfung — перевірка наявності рахунку
func (k *Kunde) HatKonto() bool {
	return k.Konto != nil
}

// die Einzahlung — поповнення рахунку
// der Betrag — сума
func (k *Konto) Einzahlen(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Einzahlung: %d", betrag)
	}
	k.Guthaben += betrag
	return nil
}

// die Abhebung — зняття коштів
func (k *Konto) Abheben(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Abhebung: %d", betrag)
	}
	if betrag > k.Guthaben {
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %d Geld abzuheben", betrag)
	}
	k.Guthaben -= betrag
	return nil
}

// die Überweisung — переказ коштів
// das Zielkonto — цільовий рахунок
func (k *Konto) Ueberweisen(ziel *Konto, betrag int) error {
	if ziel == nil {
		return fmt.Errorf("Zielkonto existiert nicht")
	}
	if err := k.Abheben(betrag); err != nil {
		return fmt.Errorf("Überweisung fehlgeschlagen: %w", err)
	}
	if err := ziel.Einzahlen(betrag); err != nil {
		// Повертаємо кошти назад у разі помилки зарахування
		k.Guthaben += betrag
		return fmt.Errorf("Rückabwicklung der Überweisung: %w", err)
	}
	return nil
}

// die Fehlerprüfung — перевірка помилок
func FehlerPruefen(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

func main() {
	konto := Konto{IBAN: "DE123456", Guthaben: 500}
	kunde := Kunde{ID: 7395, Name: "Erstes Konto", Konto: &konto}

	fmt.Printf("Hat Kunde ein Konto: %t\n", kunde.HatKonto())

	konto1 := Konto{IBAN: "DE1111", Guthaben: 500}
	konto2 := Konto{IBAN: "DE2222", Guthaben: 100}

	err := konto1.Ueberweisen(&konto2, 267)
	FehlerPruefen("Fehler bei Überweisung", err)

	fmt.Printf("Konto 1 Guthaben: %d\n", konto1.Guthaben)
	fmt.Printf("Konto 2 Guthaben: %d\n", konto2.Guthaben)
}
