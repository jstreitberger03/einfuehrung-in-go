package main

import "fmt"

func main() {
	// Übung 1: Gerade oder ungerade
	zahl := 7
	if zahl%2 == 0 {
		fmt.Printf("%d ist gerade\n", zahl)
	} else {
		fmt.Printf("%d ist ungerade\n", zahl)
	}

	// Übung 2: 7er-Reihe bis 70
	fmt.Println("\n7er-Reihe:")
	for i := 7; i <= 70; i += 7 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Übung 3: Rückwärts von 10 bis 1
	fmt.Println("\nRückwärts:")
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Übung 4: Switch für Note 1-5
	fmt.Println("\nNotenbewertung:")
	note := 3
	switch note {
	case 1:
		fmt.Println("Sehr gut")
	case 2:
		fmt.Println("Gut")
	case 3:
		fmt.Println("Befriedigend")
	case 4:
		fmt.Println("Ausreichend")
	case 5:
		fmt.Println("Nicht ausreichend")
	default:
		fmt.Println("Ungültige Note")
	}

	// Übung 5: FizzBuzz 1-100
	fmt.Println("\nFizzBuzz 1-100:")
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
