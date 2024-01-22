package main

import "fmt"

func main() {
	var notes [7]string
	var primes [7]int
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	notes[4] = "so"
	notes[5] = "la"
	notes[6] = "ti"
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7
	primes[4] = 11
	primes[5] = 13
	primes[6] = 17
	notes2 := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes2 := [7]int{2, 3, 5, 7, 11, 13, 17}
	strings := [3]string{
		"This is a series",
		"of long strings which",
		"would look awkward if place in one line",
	}
	for i := 0; i < len(notes); i++ {
		fmt.Printf("%d. %s %d\n", i, notes[i], primes[i])

	}
	fmt.Println(notes2)
	fmt.Println(primes2)
	fmt.Println(strings)
}
