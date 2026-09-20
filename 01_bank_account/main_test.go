package main

import (
	"testing"
)

// der Test — тест
func TestKontoOperationen(t *testing.T) {

	// die Einzahlung — поповнення
	t.Run("ErfolgreicheEinzahlung", func(t *testing.T) {
		konto := Konto{Guthaben: 100.0} // das Konto — рахунок, das Guthaben — баланс
		err := konto.GeldEinzahlen(50.0)

		if err != nil {
			t.Errorf("Unerwarteter Fehler bei Einzahlung: %v", err)
		}
		if konto.Guthaben != 150.0 {
			t.Errorf("Falsches Guthaben nach Einzahlung: got %.2f, want 150.00", konto.Guthaben)
		}
	})

	// der Betrag — сума
	t.Run("FehlerUngueltigerEinzahlbetrag", func(t *testing.T) {
		konto := Konto{Guthaben: 100.0}
		err := konto.GeldEinzahlen(-10.0)

		if err == nil {
			t.Error("Erwarteter Fehler bei negativem Betrag, aber nil erhalten")
		}
	})

	// die Abhebung — зняття
	t.Run("ErfolgreicheAbhebung", func(t *testing.T) {
		konto := Konto{Guthaben: 100.0}
		err := konto.GeldAbheben(40.0)

		if err != nil {
			t.Errorf("Unerwarteter Fehler bei Abhebung: %v", err)
		}
		if konto.Guthaben != 60.0 {
			t.Errorf("Falsches Guthaben nach Abhebung: got %.2f, want 60.00", konto.Guthaben)
		}
	})

	// ueberziehen — перевищувати (баланс)
	t.Run("FehlerGuthabenUeberzogen", func(t *testing.T) {
		konto := Konto{Guthaben: 100.0}
		err := konto.GeldAbheben(200.0)

		if err == nil {
			t.Error("Erwarteter Fehler bei Abhebung von zu viel Geld, aber nil erhalten")
		}
	})
}
