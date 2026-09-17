package bank

import "fmt"

// das Konto - банківський акаунт
type Konto struct {
	Guthaben float64 // das Guthaben - баланс
}

// das Geld - гроші
// die Geldabhebung - зняття грошей
func (k *Konto) GeldAbheben(betrag float64) error { // der Betrag - сума
	if betrag <= 0 {
		return fmt.Errorf("Fehler: der Betrag muss größer als 0 sein") // der Fehler - помилка
	} else if betrag > k.Guthaben {
		return fmt.Errorf("Fehler: nicht genug Guthaben auf dem Konto") // der Fehler - помилка
	}
	k.Guthaben -= betrag
	return nil
}
