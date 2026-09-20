package main

import "testing"

// das Testergebnis — результат тесту
// die Erwartung — очікування
// der Fehler — помилка
func TestKontoOperations(t *testing.T) {
	tests := []struct {
		name         string
		initialVal   int
		depositVal   int
		withdrawVal  int
		wantGuthaben int
		wantErr      bool
	}{
		{
			name:         "Gültige Einzahlung und Abhebung",
			initialVal:   500,
			depositVal:   200,
			withdrawVal:  100,
			wantGuthaben: 600,
			wantErr:      false,
		},
		{
			name:         "Ungültige Abhebung (nicht genug Geld)",
			initialVal:   100,
			depositVal:   0,
			withdrawVal:  500,
			wantGuthaben: 100,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Konto{IBAN: "DE123456", Guthaben: tt.initialVal}

			if tt.depositVal > 0 {
				if err := k.Einzahlen(tt.depositVal); err != nil {
					t.Fatalf("Einzahlen() unerwarteter Fehler = %v", err)
				}
			}

			err := k.Abheben(tt.withdrawVal)

			if (err != nil) != tt.wantErr {
				t.Errorf("Abheben() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}

			if k.Guthaben != tt.wantGuthaben {
				t.Errorf("Guthaben = %d, erwartet Guthaben = %d", k.Guthaben, tt.wantGuthaben)
			}
		})
	}
}

// die Überweisung — переказ коштів
// das Zielkonto — цільовий рахунок
func TestUeberweisen(t *testing.T) {
	tests := []struct {
		name              string
		startQuell        int
		startZiel         int
		betrag            int
		wantQuellGuthaben int
		wantZielGuthaben  int
		wantErr           bool
	}{
		{
			name:              "Erfolgreiche Überweisung",
			startQuell:        500,
			startZiel:         100,
			betrag:            200,
			wantQuellGuthaben: 300,
			wantZielGuthaben:  300,
			wantErr:           false,
		},
		{
			name:              "Fehlgeschlagene Überweisung (zu wenig Geld)",
			startQuell:        100,
			startZiel:         100,
			betrag:            500,
			wantQuellGuthaben: 100,
			wantZielGuthaben:  100,
			wantErr:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quellKonto := &Konto{IBAN: "DE1111", Guthaben: tt.startQuell}
			zielKonto := &Konto{IBAN: "DE2222", Guthaben: tt.startZiel}

			err := quellKonto.Ueberweisen(zielKonto, tt.betrag)

			if (err != nil) != tt.wantErr {
				t.Errorf("Ueberweisen() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}

			if quellKonto.Guthaben != tt.wantQuellGuthaben {
				t.Errorf("Quellkonto Guthaben = %d, erwartet = %d", quellKonto.Guthaben, tt.wantQuellGuthaben)
			}

			if zielKonto.Guthaben != tt.wantZielGuthaben {
				t.Errorf("Zielkonto Guthaben = %d, erwartet = %d", zielKonto.Guthaben, tt.wantZielGuthaben)
			}
		})
	}
}
