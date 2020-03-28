package main

import "fmt"

func equalizeArray(arr []int) int {
	numbers := make(map[int]int)
	for _, number := range arr {
		_, ok := numbers[number]
		if !ok {
			numbers[number] = 1
		} else {
			numbers[number] += 1
		}
	}

	// loop through our constructed enumeration map
	// 	to get the most occurring number
	// could also use math.Max(floats...)
	mostFrequent := 0
	for _, frequency := range numbers {
		if frequency > mostFrequent {
			mostFrequent = frequency
		}
	}

	// simplified - just delete all but the most occurring number
	fmt.Println(len(arr) - mostFrequent)
	return len(arr) - mostFrequent
}

func main() {
	arr1 := []int{1, 2, 2, 3}
	equalizeArray(arr1)

	arr2 := []int{1, 2, 2, 3, 4, 5}
	equalizeArray(arr2)

	arr3 := []int{3, 3, 2, 1, 3}
	equalizeArray(arr3)

	arr4 := []int{3, 3, 2, 2, 3, 1, 1, 3}
	equalizeArray(arr4)

	arr5 := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3}
	equalizeArray(arr5)
}
