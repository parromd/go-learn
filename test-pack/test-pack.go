package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"os"

	"go-go-go/sorts"
)

func main() {
	srand := rand.New(rand.NewSource(time.Now().UnixNano()))
	size, _ := strconv.Atoi(os.Args[1])
	unsorted := make([]int, size)

	for i := 0; i < size; i++ {
		unsorted[i] = srand.Intn(size)
	}
	fmt.Println(unsorted)
	sorter := sorts.New(size)
	sorter.Set(unsorted)

	// Do sorts
	bubble_comp, bubble_time := sorter.BubbleSort()
	fmt.Printf("Bubble sort: %d ns, doing %d comparisons\n", bubble_time.Nanoseconds(), bubble_comp)

	insert_comp, insert_time := sorter.InsertionSort()
	fmt.Printf("Insertion sort: %d ns, doing %d comparisons\n", insert_time.Nanoseconds(), insert_comp)

	select_comp, select_time := sorter.SelectionSort()
	fmt.Printf("Selection sort: %d ns, doing %d comparisons\n", select_time.Nanoseconds(), select_comp)

}
