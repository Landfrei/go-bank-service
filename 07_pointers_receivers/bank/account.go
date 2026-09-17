package bank

import (
	"fmt"
)

// das Konto - банківський рахунок
type Konto struct {
	Inhaber  string  // der Inhaber - власник
	Guthaben float64 // das Guthaben - баланс
}

// die Einzahlung nach Wert - поповнення через копію.
// die Einzahlung - поповнення
// der Wert - копія
func (k Konto) EinzahlenWert(betrag float64) { // der Betrag - сума
	k.Guthaben += betrag
	fmt.Printf("[Wert] Neues Guthaben fuer %s: %.2f EUR\n", k.Inhaber, k.Guthaben)
}

// der Zeiger - вказівник
func (k *Konto) EinzahlenZeiger(betrag float64) { // der Betrag - сума
	k.Guthaben += betrag
	fmt.Printf("[Zeiger] Neues Guthaben fuer %s: %.2f EUR\n", k.Inhaber, k.Guthaben)
}
