package bank

import (
	"testing"
)

func TestAbwickeln(t *testing.T) {
	tests := []struct {
		Name        string // der Name — назва
		Initial     int    // das Initial — початкове значення
		Amount      int    // der Betrag — сума
		WantErr     bool   // der Fehler — помилка
		ExpectedBal int    // der Saldo — баланс
	}{
		{
			Name:        "ErfolgreicheAbhebung", // die erfolgreiche Abhebung — успішне зняття коштів
			Initial:     100,
			Amount:      40,
			WantErr:     false,
			ExpectedBal: 60,
		},
		{
			Name:        "FehlerGuthabenUeberzogen", // der Fehler: das Guthaben überzogen — помилка: перевищення балансу
			Initial:     100,
			Amount:      150,
			WantErr:     true,
			ExpectedBal: 100,
		},
		{
			Name:        "FehlerUngueltigerBetrag", // der Fehler: der ungültige Betrag — помилка: недійсна сума
			Initial:     100,
			Amount:      0,
			WantErr:     true,
			ExpectedBal: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// das Muster — зразок / шаблон
			// Max Mustermann — Макс Шаблоненко
			konto := Konto{Inhaber: "Max Mustermann", Guthaben: tt.Initial} // das Konto — банківський рахунок
			err := Abwickeln(&konto, tt.Amount)

			if (err != nil) != tt.WantErr {
				t.Errorf("Abwickeln() error = %v, wantErr %v", err, tt.WantErr)
			}

			if konto.Guthaben != tt.ExpectedBal {
				t.Errorf("Guthaben = %d, expected %d", konto.Guthaben, tt.ExpectedBal)
			}
		})
	}
}
