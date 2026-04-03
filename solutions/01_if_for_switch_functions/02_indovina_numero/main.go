package main

import "fmt"

func Confronta(tentativo, segreto int) string {
	if tentativo > segreto {
		return "Troppo alto"
	}
	if tentativo < segreto {
		return "Troppo basso"
	}
	return "Indovinato!"
}

func ContaTentativi(storico []int) int {
	return len(storico)
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
