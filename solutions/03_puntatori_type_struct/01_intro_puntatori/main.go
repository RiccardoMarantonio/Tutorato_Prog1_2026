package main

import "fmt"

func Scambia(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func StampaInfo(valore int, puntatore *int) {
	fmt.Printf("Valore: %d, Indirizzo: %p, Valore dereferenziato: %d\n", valore, puntatore, *puntatore)
}

func main() {
	x := 5
	y := 10
	fmt.Printf("Prima dello scambio: x = %d, y = %d\n", x, y)
	fmt.Printf("Indirizzo di x: %p\n", &x)
	fmt.Printf("Indirizzo di y: %p\n", &y)

	Scambia(&x, &y)

	fmt.Printf("Dopo lo scambio: x = %d, y = %d\n", x, y)
	StampaInfo(x, &x)
}
