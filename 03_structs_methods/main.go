package main

import "fmt"

// das Konto - рахунок

type Konto struct {
	IBAN     string
	Guthaben int // das Guthaben - баланс, залишок на рахунку
}

// der Kunde - клієнт
type Kunde struct {
	ID    int
	Name  string
	Konto *Konto
}

func (k *Kunde) HatKonto() bool {
	if k.Konto != nil {
		return true
	}
	return false
}

// einzahlen - поповнювати
// der Betrag - сума
func (k *Konto) Einzahlen(betrag int) {
	if betrag > 0 {
		k.Guthaben += betrag
	}
}

// abheben - знімати
func (k *Konto) Abheben(betrag int) bool {
	if betrag > k.Guthaben {
		return false
	}
	k.Guthaben -= betrag
	return true
}

// überweisen = переказувати(гроші)
func (k *Konto) Ueberweisen(ziel *Konto, betrag int) bool {
	if k.Abheben(betrag) {
		ziel.Einzahlen(betrag)
		return true
	}
	return false
}

func main() {
	konto := Konto{"DE123456", 500}
	kunde := Kunde{7395, "Erstes Konto", &konto}

	fmt.Println(kunde.HatKonto())

	konto1 := Konto{"DE1111", 500}
	konto2 := Konto{"DE2222", 100}
	//der Erfolg = успіх
	erfolg := konto1.Ueberweisen(&konto2, 267)
	//erfolgreich = успішний
	fmt.Println("Ueberweisung erfolgreich:", erfolg)
	fmt.Println("Konto 1 Guthaben:", konto1.Guthaben)
	fmt.Println("Konto 2 Guthaben:", konto2.Guthaben)
}
