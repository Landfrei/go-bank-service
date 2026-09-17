package bank

// die Schnittstelle - інтерфейс
// das Zahlungsmittel - платіжний засіб
type Zahlungsmittel interface {
	Bezahlen(betrag float64) bool // der Betrag - сума, die Bezahlung - оплата
}

// die Kreditkarte - кредитна картка
type Kreditkarte struct {
	Limit float64 // das Limit - ліміт
}

// die Bezahlung - оплата, der Betrag - сума
func (k *Kreditkarte) Bezahlen(betrag float64) bool {
	if betrag <= k.Limit {
		k.Limit -= betrag
		return true
	}
	return false
}

// das Konto - рахунок
type Konto struct {
	Guthaben float64 // das Guthaben - баланс
}

// die Bezahlung - оплата, der Betrag - сума
func (k *Konto) Bezahlen(betrag float64) bool {
	if betrag <= k.Guthaben {
		k.Guthaben -= betrag
		return true
	}
	return false
}
