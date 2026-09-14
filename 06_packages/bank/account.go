package bank

import "fmt"

type Konto struct {
	Owner    string
	Guthaben int
}

type GuthabenUeberzigenError struct {
	Fehlenderbetrag int
}

func (g *GuthabenUeberzigenError) Error() string {
	return fmt.Sprintf("Fehler! Es fehlen %d EUR", g.Fehlenderbetrag)
}

func Abwickeln(k *Konto, betrag int) error {
	if betrag > k.Guthaben {
		return &GuthabenUeberzigenError{Fehlenderbetrag: betrag - k.Guthaben}
	}
	k.Guthaben -= betrag
	return nil
}
