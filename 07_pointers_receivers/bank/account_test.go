package bank

import (
	"testing"
)

func TestEinzahlen(t *testing.T) {
	// das Muster — зразок / шаблон
	// Max Mustermann — Макс Шаблоненко

	t.Run("EinzahlenWertKeineAenderung", func(t *testing.T) { // der Testfall — тестовий випадок
		konto := Konto{Inhaber: "Max Mustermann", Guthaben: 100.0} // das Konto — банківський рахунок
		konto.EinzahlenWert(50.0)

		if konto.Guthaben != 100.0 {
			// das Guthaben — баланс
			t.Errorf("EinzahlenWert hat Guthaben geaendert: got %f, want 100.0", konto.Guthaben)
		}
	})

	t.Run("EinzahlenZeigerAenderung", func(t *testing.T) {
		konto := Konto{Inhaber: "Max Mustermann", Guthaben: 100.0} // das Konto — банківський рахунок
		konto.EinzahlenZeiger(50.0)

		if konto.Guthaben != 150.0 {
			// das Guthaben — баланс
			t.Errorf("EinzahlenZeiger hat Guthaben nicht geaendert: got %f, want 150.0", konto.Guthaben)
		}
	})
}
