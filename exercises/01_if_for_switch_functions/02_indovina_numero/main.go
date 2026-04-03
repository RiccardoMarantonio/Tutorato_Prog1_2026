package main

import "fmt"

// Confronta restituisce un feedback sul tentativo rispetto al segreto.
func Confronta(tentativo, segreto int) string {
	// TODO: implementare
	return ""
}

// ContaTentativi restituisce il numero di tentativi effettuati.
func ContaTentativi(storico []int) int {
	// TODO: implementare
	return 0
}

func main() {
	segreto := 42
	var storico []int
	for {
		var tentativo int
		fmt.Scan(&tentativo)
		if tentativo < 1 || tentativo > 100 {
			fmt.Println("Inserisci un numero tra 1 e 100")
			continue
		}
		storico = append(storico, tentativo)
		feedback := Confronta(tentativo, segreto)
		fmt.Println(feedback)
		if feedback == "Indovinato!" {
			fmt.Printf("Hai indovinato in %d tentativi!\n", ContaTentativi(storico))
			break
		}
	}
}
