package main

import "fmt"

func ConfrontaParola(input, segreto string) string {
	if input == segreto {
		return "Hai indovinato!"
	}
	if len(input) < len(segreto) {
		return "Troppo corto"
	}
	if len(input) > len(segreto) {
		return "Troppo lungo"
	}
	return "Lunghezza giusta, riprova!"
}

func main() {
	segreto := "golang"
	maxTentativi := 6
	vinto := false

	for i := 1; i <= maxTentativi; i++ {
		var tentativo string
		fmt.Scan(&tentativo)
		feedback := ConfrontaParola(tentativo, segreto)
		fmt.Printf("Tentativo %d/%d: %s\n", i, maxTentativi, feedback)
		if feedback == "Hai indovinato!" {
			fmt.Printf("Vittoria in %d tentativi!\n", i)
			vinto = true
			break
		}
	}

	if !vinto {
		fmt.Printf("Hai perso! La parola era: %s\n", segreto)
	}
}
