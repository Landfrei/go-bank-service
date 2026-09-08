package main

import (
	"fmt"
)

// das Konto = обліковий запис
type Konto struct {
	Guthaben int
}

// die Einzahlung = поповнення грошами
func (k *Konto) GeldEinzahlen(betrag int) {
	k.Guthaben += betrag
}

// der Betrag = сума

// die Abhebung = зняття грошей
func (k *Konto) GeldAbheben(betrag int) error {
	if betrag > k.Guthaben {
		return fmt.Errorf("nicht genug Geld auf dem Konto, um %d Geld abzuheben", betrag)
	}
	k.Guthaben -= betrag
	return nil

}

func main() {
	meinBank := Konto{
		Guthaben: 100,
	}

	meinBank.GeldEinzahlen(50)
	fmt.Println("Mein Guthaben ist ", meinBank.Guthaben, "Geld")

	err := meinBank.GeldAbheben(200)
	if err != nil {
		fmt.Println(err)
	}

}
