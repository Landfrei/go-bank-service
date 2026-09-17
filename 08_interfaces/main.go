package main

import (
	"08_interfaces/bank"
	"fmt"
)

// die Bezahlung - оплата
func EinkaufMachen(z bank.Zahlungsmittel, betrag float64) { // der Betrag - сумма
	erfolgreich := z.Bezahlen(betrag)
	if erfolgreich {
		fmt.Printf("Bezahlung von %.2f EUR erfolgreich!\n", betrag)
	} else {
		fmt.Printf("Bezahlung von %.2f EUR fehlgeschlagen!\n", betrag)
	}
}
func main() {
	kreditkarte := &bank.Kreditkarte{Limit: 1000.0}
	konto := &bank.Konto{Guthaben: 500.0}
	// der Einkauf - покупкa
	EinkaufMachen(kreditkarte, 200.0)
	EinkaufMachen(konto, 600.0)
}
