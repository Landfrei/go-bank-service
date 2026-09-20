package main

import (
	"errors"
	"testing"
)

func TestAbwickelnErrors(t *testing.T) {
	tests := []struct {
		Name              string // der Name — назва тесту
		Guthaben          int    // das Guthaben — баланс / власні кошти
		Betrag            int    // der Betrag — сума
		ErwarteteGuthaben int    // das erwartete Guthaben — очікуваний баланс
		HatFehler         bool   // der Fehler — наявність помилки
		Fehlbetrag        int    // der Fehlbetrag — сума, якої бракує
	}{
		{
			Name:              "Erfolgreiche Abhebung", // erfolgreiche Abhebung — успішне зняття коштів
			Guthaben:          100,
			Betrag:            40,
			ErwarteteGuthaben: 60,
			HatFehler:         false,
			Fehlbetrag:        0,
		},
		{
			Name:              "Fehler: Guthaben ueberzogen", // das Guthaben überzogen — баланс перевищено
			Guthaben:          60,
			Betrag:            90,
			ErwarteteGuthaben: 60,
			HatFehler:         true,
			Fehlbetrag:        30,
		},
		{
			Name:              "Fehler: Ungueltiger Betrag", // ungültiger Betrag — некоректна сума
			Guthaben:          100,
			Betrag:            -10,
			ErwarteteGuthaben: 100,
			HatFehler:         true,
			Fehlbetrag:        0,
		},
	}

	for _, tt := range tests {
		// ausführen — виконувати
		t.Run(tt.Name, func(t *testing.T) {
			gotGuthaben, err := Abwickeln(tt.Guthaben, tt.Betrag)

			// die Fehlerprüfung — перевірка помилки
			if (err != nil) != tt.HatFehler {
				t.Fatalf("Abwickeln() Fehler = %v, erwartet Fehler = %v", err, tt.HatFehler)
			}

			// der Fehlertyp — тип помилки
			if tt.Fehlbetrag > 0 {
				var customErr *GuthabenUeberzogenFehler
				if errors.As(err, &customErr) {
					if customErr.Fehlbetrag != tt.Fehlbetrag {
						t.Errorf("Fehlbetrag = %d, erwartet = %d", customErr.Fehlbetrag, tt.Fehlbetrag)
					}
				} else {
					// unerwartet — неочікуваний
					t.Errorf("Unerwarteter Fehlertyp: %T", err)
				}
			}

			// der Vergleich — порівняння / перевірка результату
			if gotGuthaben != tt.ErwarteteGuthaben {
				t.Errorf("Guthaben = %d, erwartet Guthaben = %d", gotGuthaben, tt.ErwarteteGuthaben)
			}
		})
	}
}
