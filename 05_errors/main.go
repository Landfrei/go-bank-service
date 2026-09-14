package main

import "fmt"

type GuthabenUeberzogenError struct {
	FehlenderBetrag int
}

func (g *GuthabenUeberzogenError) Error() string {
	return fmt.Sprintf("Konto ueberzogen! Es fehlen %d EUR", g.FehlenderBetrag)
}

func Abwickeln(guthaben int, betrag int) (int, error) {
	if betrag > guthaben {
		return guthaben, &GuthabenUeberzogenError{FehlenderBetrag: betrag - guthaben}
	}
	return guthaben - betrag, nil
}

func main() {
	guthaben := 100

	neuesGuthaben, err := Abwickeln(guthaben, 40)

	if err != nil {
		fmt.Println("Fehler: ", err)
	} else {
		guthaben = neuesGuthaben
	}
	neuesGuthaben, err = Abwickeln(guthaben, 90)

	if err != nil {
		fmt.Println("Fehler: ", err)
	} else {
		guthaben = neuesGuthaben
	}
	fmt.Println("Deines Guthaben: ", guthaben)

}
