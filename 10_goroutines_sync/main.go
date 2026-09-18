package main

import (
	"fmt"
	"sync"
	"time"
)

// die Bearbeitung - опрацювання замовлення
func Bearbeitung(kassenID int, zaehler *sync.WaitGroup) { // die Kasse - каса, der Zaehler - лічильник, die Wartegruppe - група очікування
	defer zaehler.Done()
	fmt.Printf("[die Kasse %d] Beginn der Bearbeitung der Bestellung\n", kassenID) // Beginn der Bearbeitung der Bestellung - починає опрацювання замовлення
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[die Kasse %d] Die Bestellung wurde erfolgreich bearbeitet! \n", kassenID) // die Bestellung wurde erfolgreich bearbeitet - замовлення успішно опрацьоване
}

func main() {
	var zaehler sync.WaitGroup // der Zaehler - лічильник, die Wartegruppe - група очікування
	gesamtenBestellungen := 3  // die gesamten Bestellungen - загальна кількість замовлень
	fmt.Println(" ===== START DES GESCHÄFTS ===== ")

	for i := 1; i <= gesamtenBestellungen; i++ {
		zaehler.Add(1)
		go Bearbeitung(i, &zaehler)
	}
	fmt.Println("Die Hauptkasse wartet auf das Ende aller Aufgaben...") // Die Hauptkasse wartet auf das Ende aller Aufgaben - головна каса очікує завершення всіх завдань
	zaehler.Wait()
	fmt.Println(" ===== ALLE BESTELLUNGEN WURDEN ERFOLGREICH BEARBEITET! ===== ") // Alle Bestellungen wurden erfolgreich bearbeitet - всі замовлення успішно опрацьовані
}
