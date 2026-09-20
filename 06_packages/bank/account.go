package bank

import (
	"fmt"
)

// das Konto — банківський рахунок
type Konto struct {
	Inhaber  string // der Inhaber — власник рахунку
	Guthaben int    // das Guthaben — баланс / власні кошти
}

// der GuthabenUeberzogenFehler — помилка перевищення балансу
type GuthabenUeberzogenFehler struct {
	Fehlbetrag int // der Fehlbetrag — сума, якої бракує
}

func (g *GuthabenUeberzogenFehler) Error() string {
	// fehlen — бракувати / не вистачати
	return fmt.Sprintf("Konto ueberzogen! Es fehlen %d EUR", g.Fehlbetrag)
}

// die Abwicklung — проведення платежу
// der Betrag — сума
func Abwickeln(k *Konto, betrag int) error {
	if betrag <= 0 {
		// ungueltig — недійсний / некоректний
		return fmt.Errorf("ungueltiger Betrag: %d EUR", betrag)
	}
	if betrag > k.Guthaben {
		return &GuthabenUeberzogenFehler{Fehlbetrag: betrag - k.Guthaben}
	}
	// abziehen — віднімати
	k.Guthaben -= betrag
	return nil
}
