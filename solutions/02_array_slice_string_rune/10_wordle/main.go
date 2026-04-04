package main

import "fmt"

func Feedback(tentativo, segreto string) []rune {
	t := []rune(tentativo)
	s := []rune(segreto)
	n := len(t)
	result := make([]rune, n)

	// Conta le occorrenze nel segreto
	countSegreto := make(map[rune]int)
	for _, r := range s {
		countSegreto[r]++
	}

	// Prima passata: verdi
	for i := 0; i < n; i++ {
		if t[i] == s[i] {
			result[i] = 'V'
			countSegreto[t[i]]--
		}
	}

	// Seconda passata: gialli e grigi
	for i := 0; i < n; i++ {
		if result[i] == 'V' {
			continue
		}
		if countSegreto[t[i]] > 0 {
			result[i] = 'G'
			countSegreto[t[i]]--
		} else {
			result[i] = 'X'
		}
	}

	return result
}

func StampaFeedback(feedback []rune) {
	for i, r := range feedback {
		if i > 0 {
			fmt.Print(" ")
		}
		switch r {
		case 'V':
			fmt.Print("VERDE")
		case 'G':
			fmt.Print("GIALLO")
		case 'X':
			fmt.Print("GRIGIO")
		}
	}
	fmt.Println()
}

func main() {
	segreto := "crane"
	maxTentativi := 6
	vinto := false

	for i := 1; i <= maxTentativi; i++ {
		var tentativo string
		fmt.Scan(&tentativo)

		if len([]rune(tentativo)) != 5 {
			fmt.Println("La parola deve avere 5 lettere")
			i--
			continue
		}

		feedback := Feedback(tentativo, segreto)
		fmt.Printf("Tentativo %d/%d: %s\n", i, maxTentativi, tentativo)
		StampaFeedback(feedback)

		tuttiVerdi := true
		for _, r := range feedback {
			if r != 'V' {
				tuttiVerdi = false
				break
			}
		}
		if tuttiVerdi {
			fmt.Printf("Hai indovinato! Vittoria in %d tentativi!\n", i)
			vinto = true
			break
		}
	}

	if !vinto {
		fmt.Printf("Hai perso! La parola era: %s\n", segreto)
	}
}
