package main

import "testing"

// das Testergebnis - результат тесту
// die Erwartung - очікування

func TestKontoOperations(t *testing.T) {
	tests := []struct {
		name         string
		initialVal   int
		depositVal   int
		withdrawVal  int
		wantGuthaben int
		wantSuccess  bool
	}{
		{
			name:         "Успішне поповнення та зняття",
			initialVal:   500,
			depositVal:   200,
			withdrawVal:  100,
			wantGuthaben: 600,
			wantSuccess:  true,
		},
		{
			name:         "Спроба зняти більше, ніж є",
			initialVal:   100,
			depositVal:   0,
			withdrawVal:  500,
			wantGuthaben: 100,
			wantSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Konto{IBAN: "DE123456", Guthaben: tt.initialVal}

			if tt.depositVal > 0 {
				k.Einzahlen(tt.depositVal)
			}

			gotSuccess := k.Abheben(tt.withdrawVal)

			if gotSuccess != tt.wantSuccess {
				t.Errorf("Abheben() success = %v, want %v", gotSuccess, tt.wantSuccess)
			}

			if k.Guthaben != tt.wantGuthaben {
				t.Errorf("Guthaben = %d, want %d", k.Guthaben, tt.wantGuthaben)
			}
		})
	}
}
