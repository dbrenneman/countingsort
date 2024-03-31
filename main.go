package main

import (
	"fmt"
	"math/rand"
)

type Customer struct {
	id           string
	numPurchases int
}

func makeRandomSlice(numItems, max int) []Customer {
	s := make([]Customer, numItems)

	for i := 0; i < numItems; i++ {
		s[i] = Customer{
			id:           fmt.Sprintf("C%d", i),
			numPurchases: rand.Intn(max),
		}
	}

	return s
}

func printSlice(s []Customer, max int) {
	for i, v := range s {
		if i < max {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func checkSorted(s []Customer) {
	sorted := true
	for i := 1; i < len(s); i++ {
		if s[i-1].numPurchases > s[i].numPurchases {
			sorted = false
		}
	}

	if sorted {
		fmt.Println("The slice is sorted.")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
}

func countingSort(s []Customer, max int) []Customer {
	count := make([]int, max)

	out := make([]Customer, len(s))

	for _, v := range s {
		count[v.numPurchases]++
	}

	for i, v := range count {
		if i > 0 {
			count[i] = v + count[i-1]
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		num := s[i].numPurchases
		out[count[num]-1] = s[i]
		count[num] -= 1
	}

	return out
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
