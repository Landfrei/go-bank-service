package main

import (
	"fmt"
)

// der GuthabenUeberzogenFehler — помилка перевищення балансу
type GuthabenUeberzogenFehler struct {
	// der Fehlbetrag — сума, якої бракує
	Fehlbetrag int
}

func (g *GuthabenUeberzogenFehler) Error() string {
	// das Konto — банківський рахунок
	// fehlen — бракувати / не вистачати
	return fmt.Sprintf("Konto ueberzogen! Es fehlen %d EUR", g.Fehlbetrag)
}

// die Abwicklung — проведення платежу
// das Guthaben — баланс / власні кошти
// der Betrag — сума
func Abwickeln(guthaben int, betrag int) (int, error) {
	if betrag <= 0 {
		// ungueltig — недійсний / некоректний
		return guthaben, fmt.Errorf("ungueltiger Betrag: %d EUR", betrag)
	}
	if betrag > guthaben {
		return guthaben, &GuthabenUeberzogenFehler{Fehlbetrag: betrag - guthaben}
	}
	return guthaben - betrag, nil
}

// die Fehlerprüfung — перевірка помилки
func FehlerPruefen(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

func main() {
	guthaben := 100
	var err error

	guthaben, err = Abwickeln(guthaben, 40)
	FehlerPruefen("Fehler bei Abwicklung 1", err)

	guthaben, err = Abwickeln(guthaben, 90)
	FehlerPruefen("Fehler bei Abwicklung 2", err)

	fmt.Printf("Deines Guthaben: %d EUR\n", guthaben)
}
