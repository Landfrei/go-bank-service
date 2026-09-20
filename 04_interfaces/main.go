package main

import (
	"fmt"
)

// die Zahlungsmethode — інтерфейс для способів оплати
type Zahlungsmethode interface {
	Zahlen(betrag int) error
}

// die Kreditkarte — кредитна картка
// das Limit — кредитний ліміт
// das Guthaben — власні кошти на картці
type Kreditkarte struct {
	Limit    int
	Guthaben int
}

// die Zahlung — здійснення платежу карткою
func (k *Kreditkarte) Zahlen(betrag int) error {
	if betrag <= 0 {
		return fmt.Errorf("ungültiger Zahlungsbetrag: %d", betrag)
	}
	if betrag > k.Guthaben+k.Limit {
		return fmt.Errorf("Kreditlimit überschritten: %d Geld gefordert, verfügbares Limit + Guthaben = %d", betrag, k.Guthaben+k.Limit)
	}
	k.Guthaben -= betrag
	return nil
}

// die Abwicklung — обробка платежу через будь-який спосіб оплати
func Abwickeln(z Zahlungsmethode, betrag int) error {
	if err := z.Zahlen(betrag); err != nil {
		return fmt.Errorf("Abwicklung fehlgeschlagen: %w", err)
	}
	fmt.Printf("Zahlung von %d Geld erfolgreich verarbeitet.\n", betrag)
	return nil
}

// die Fehlerprüfung — перевірка помилок
func FehlerPruefen(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

func main() {
	karte := Kreditkarte{Limit: 3660, Guthaben: 150}

	err := Abwickeln(&karte, 250)
	FehlerPruefen("Fehler bei Abwicklung 1", err)
	fmt.Printf("Kartenstatus: Guthaben = %d, Limit = %d\n", karte.Guthaben, karte.Limit)

	err = Abwickeln(&karte, 4000) // Перевищує ліміт + баланс
	FehlerPruefen("Fehler bei Abwicklung 2", err)
	fmt.Printf("Kartenstatus: Guthaben = %d, Limit = %d\n", karte.Guthaben, karte.Limit)
}
