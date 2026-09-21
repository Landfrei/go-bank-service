package bank

import (
	"testing"
)

func TestZahlungsmittel(t *testing.T) {
	tests := []struct {
		Name     string         // der Name - назва
		Zahlung  Zahlungsmittel // das Zahlungsmittel - платіжний засіб
		Betrag   float64        // der Betrag - сума
		Erwartet bool           // das erwartete Ergebnis - очікуваний результат
		RestWert float64        // der Restwert - залишок
	}{
		{
			Name:     "Kreditkarte erfolgreich", // Kreditkarte erfolgreich - кредитна картка успішно
			Zahlung:  &Kreditkarte{Limit: 500.0},
			Betrag:   200.0,
			Erwartet: true,
			RestWert: 300.0,
		},
		{
			Name:     "Kreditkarte abgelehnt", // Kreditkarte abgelehnt - кредитна картка відхилена
			Zahlung:  &Kreditkarte{Limit: 100.0},
			Betrag:   200.0,
			Erwartet: false,
			RestWert: 100.0,
		},
		{
			Name:     "Konto erfolgreich", // Konto erfolgreich - рахунок успішно
			Zahlung:  &Konto{Guthaben: 1000.0},
			Betrag:   450.0,
			Erwartet: true,
			RestWert: 550.0,
		},
		{
			Name:     "Konto abgelehnt", // Konto abgelehnt - рахунок відхилений
			Zahlung:  &Konto{Guthaben: 50.0},
			Betrag:   100.0,
			Erwartet: false,
			RestWert: 50.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ergebnis := tt.Zahlung.Bezahlen(tt.Betrag)
			if ergebnis != tt.Erwartet {
				t.Errorf("Bezahlen() = %v, erwartet %v", ergebnis, tt.Erwartet)
			}

			switch z := tt.Zahlung.(type) {
			case *Kreditkarte:
				if z.Limit != tt.RestWert {
					t.Errorf("Kreditkarte Limit = %v, erwartet %v", z.Limit, tt.RestWert)
				}
			case *Konto:
				if z.Guthaben != tt.RestWert {
					t.Errorf("Konto Guthaben = %v, erwartet %v", z.Guthaben, tt.RestWert)
				}
			}
		})
	}
}
