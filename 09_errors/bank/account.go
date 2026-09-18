package bank

import "fmt"

// das Konto - банківський акаунт
type Konto struct {
	Guthaben float64 // das Guthaben - баланс
}

// unzureichend - недостатній
// das Guthaben - баланс
// der Fehler - помилка
//die UnzureichendesGuthabenFehler - помилка недостатнього балансу
type UnzureichendesGuthabenFehler struct {
	Guthaben float64
	Betrag   float64
}

func (e *UnzureichendesGuthabenFehler) Error() string {
	return fmt.Sprintf("nicht genug Guthaben auf dem Konto: %.2f gefordert, aber nur %.2f vorhanden", e.Betrag, e.Guthaben) // der Fehler - помилка, gefordert - вимагався, vorhanden - наявний
}
func (k *Konto) GeldAbheben(betrag float64) error { // der Betrag - сума
	if betrag <= 0 {
		return fmt.Errorf("Fehler: der Betrag muss größer als 0 sein") // der Fehler - помилка
	} else if betrag > k.Guthaben {
		return &UnzureichendesGuthabenFehler{
			Guthaben: k.Guthaben,
			Betrag:   betrag,
		}
	}
	k.Guthaben -= betrag
	return nil
}
