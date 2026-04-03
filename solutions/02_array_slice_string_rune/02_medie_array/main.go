package main

import "fmt"

func MediaArray(arr [10]int) float64 {
	somma := 0
	for _, v := range arr {
		somma += v
	}
	return float64(somma) / 10
}

func ContaSopraMedia(arr [10]int, media float64) int {
	count := 0
	for _, v := range arr {
		if float64(v) > media {
			count++
		}
	}
	return count
}

func ContaSottoMedia(arr [10]int, media float64) int {
	count := 0
	for _, v := range arr {
		if float64(v) < media {
			count++
		}
	}
	return count
}

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	media := MediaArray(arr)
	fmt.Printf("Media: %.1f\n", media)
	fmt.Printf("Sopra la media: %d\n", ContaSopraMedia(arr, media))
	fmt.Printf("Sotto la media: %d\n", ContaSottoMedia(arr, media))
}
