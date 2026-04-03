package main

import "fmt"

// Confronta restituisce un feedback sul tentativo rispetto al segreto.
func Confronta(tentativo, segreto int) string {
	// TODO: implementare
	return ""
}

func main() {
	segreto := 42
	tentativi := 0
	for {
		var tentativo int
		fmt.Scan(&tentativo)
		if tentativo < 1 || tentativo > 100 {
			fmt.Println("Inserisci un numero tra 1 e 100")
			continue
		}
		tentativi++
		feedback := Confronta(tentativo, segreto)
		fmt.Println(feedback)
		if feedback == "Indovinato!" {
			fmt.Printf("Hai indovinato in %d tentativi!\n", tentativi)
			break
		}
	}
}
