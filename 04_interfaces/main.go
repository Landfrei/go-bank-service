package main

import (
	"fmt"
)

type Payer interface {
	Pay(betrag int) bool
}

type Kreditkarte struct {
	Limit    int
	Guthaben int
}

func (k *Kreditkarte) Pay(betrag int) bool {
	if betrag > k.Guthaben+k.Limit {
		return false
	}
	k.Guthaben -= betrag
	return true
}

func Abwickeln(p Payer, betrag int) {
	erfolg := p.Pay(betrag)
	fmt.Println("Erfolg: ", erfolg)
}

func main() {
	karte := Kreditkarte{Limit: 3660, Guthaben: 150}
	Abwickeln(&karte, 250)
	fmt.Println("", karte)
	Abwickeln(&karte, 250)
	fmt.Println("", karte)
}
