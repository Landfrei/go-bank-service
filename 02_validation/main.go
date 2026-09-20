package main

import (
	"fmt"
)

// das Konto — банківський рахунок
// das Guthaben — баланс
type Konto struct {
	Guthaben int
}

// die Einzahlung — поповнення рахунку
// der Betrag — сума
func (k *Konto) GeldEinzahlen(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Einzahlung: %d", betrag)
	}
	k.Guthaben += betrag
	return nil
}

// die Abhebung — зняття коштів
// ungültig — некоректний / недійсний
func (k *Konto) GeldAbheben(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Betrag für eine Abhebung: %d", betrag)
	}
	if betrag > k.Guthaben {
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %d Geld abzuheben", betrag)
	}
	k.Guthaben -= betrag
	return nil
}

// gültig — дійсний / коректний
// das Landeskürzel — код країни
func IstIBANGueltig(iban string) error {
	if len(iban) != 22 {
		return fmt.Errorf("ungültige IBAN-Länge: %d (muss 22 sein)", len(iban))
	}
	if iban[:2] != "DE" {
		return fmt.Errorf("ungültiges Landeskürzel: %s (muss DE sein)", iban[:2])
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
	meinKonto := Konto{Guthaben: 100}

	FehlerPruefen("Fehler bei Einzahlung", meinKonto.GeldEinzahlen(-50))
	FehlerPruefen("Fehler bei Abhebung", meinKonto.GeldAbheben(200))
	FehlerPruefen("Fehler bei IBAN", IstIBANGueltig("UA1234567890123456789012"))

	fmt.Printf("Mein Guthaben ist %d Geld\n", meinKonto.Guthaben)
}
