package main

import (
	"06_packages/bank"
	"fmt"
)

func main() {
	k := bank.Konto{Owner: "Litler", Guthaben: 100}
	err := bank.Abwickeln(&k, 40)
	if err != nil {
		fmt.Println("Fehler: ", err)
	}
	fmt.Println("Guthaben: ", k.Guthaben)
	err = bank.Abwickeln(&k, 7676)
	if err != nil {
		fmt.Println("Fehler: ", err)
	}
	fmt.Println("Guthaben: ", k.Guthaben)

}
