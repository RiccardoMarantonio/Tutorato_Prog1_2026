package main

import "fmt"

func Scambia(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

type Punto struct {
	X float64
	Y float64
}

func Sqrt(x float64) float64 {
	if x < 0 {
		return 0
	}
	z := x
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func DistanzaDaOrigine(p Punto) float64 {
	return Sqrt(p.X*p.X + p.Y*p.Y)
}

func Sposta(p *Punto, dx, dy float64) {
	p.X += dx
	p.Y += dy
}

func main() {
	x := 5
	y := 10
	fmt.Println("=== Puntatori ===")
	fmt.Printf("Prima: x = %d, y = %d\n", x, y)
	Scambia(&x, &y)
	fmt.Printf("Dopo:  x = %d, y = %d\n", x, y)

	fmt.Println("\n=== Struct Punto ===")
	var px, py float64
	fmt.Scan(&px, &py)
	p := Punto{X: px, Y: py}
	fmt.Printf("Punto: (%.2f, %.2f)\n", p.X, p.Y)
	fmt.Printf("Distanza dall'origine: %.2f\n", DistanzaDaOrigine(p))

	var dx, dy float64
	fmt.Scan(&dx, &dy)
	Sposta(&p, dx, dy)
	fmt.Printf("Dopo spostamento (%.2f, %.2f): (%.2f, %.2f)\n", dx, dy, p.X, p.Y)
}
