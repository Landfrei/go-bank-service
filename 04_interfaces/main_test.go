package main

import "testing"

// das Testergebnis — результат тесту
// die Erwartung — очікування
func TestKreditkarteZahlen(t *testing.T) {
	tests := []struct {
		name         string
		limit        int
		guthaben     int
		betrag       int
		wantGuthaben int
		wantErr      bool
	}{
		{
			name:         "Erfolgreiche Zahlung mit Guthaben",
			limit:        1000,
			guthaben:     500,
			betrag:       200,
			wantGuthaben: 300,
			wantErr:      false,
		},
		{
			name:         "Erfolgreiche Zahlung mit Kreditlimit",
			limit:        1000,
			guthaben:     100,
			betrag:       500,
			wantGuthaben: -400,
			wantErr:      false,
		},
		{
			name:         "Fehlgeschlagen (Limit überschritten)",
			limit:        500,
			guthaben:     100,
			betrag:       1000,
			wantGuthaben: 100,
			wantErr:      true,
		},
		{
			name:         "Ungültiger Betrag",
			limit:        500,
			guthaben:     100,
			betrag:       -50,
			wantGuthaben: 100,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			karte := &Kreditkarte{Limit: tt.limit, Guthaben: tt.guthaben}

			err := Abwickeln(karte, tt.betrag)

			if (err != nil) != tt.wantErr {
				t.Errorf("Abwickeln() Fehler = %v, erwartet Fehler = %v", err, tt.wantErr)
			}

			if karte.Guthaben != tt.wantGuthaben {
				t.Errorf("Guthaben = %d, erwartet Guthaben = %d", karte.Guthaben, tt.wantGuthaben)
			}
		})
	}
}
