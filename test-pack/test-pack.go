package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"go-go-go/sorts"
)

func main() {
	srand := rand.New(rand.NewSource(time.Now().UnixNano()))
	size, _ := strconv.Atoi(os.Args[1])
	unsorted := make([]int, size)

	for i := 0; i < size; i++ {
		unsorted[i] = srand.Intn(size)
	}
	// fmt.Println(unsorted)
	sorter := sorts.New(size)
	sorter.Set(unsorted)
	format := "%s: %d ms, doing %d comparisons\n"

	// Do sorts
	bubble_comp, bubble_time := sorter.BubbleSort()
	fmt.Printf(format, "Bubble", bubble_time.Milliseconds(), bubble_comp)

	insert_comp, insert_time := sorter.InsertionSort()
	fmt.Printf(format, "Insertion", insert_time.Milliseconds(), insert_comp)

	select_comp, select_time := sorter.SelectionSort()
	fmt.Printf(format, "Selection", select_time.Milliseconds(), select_comp)

	merge_comp, merge_time := sorter.MergeSort()
	fmt.Printf(format, "Merge", merge_time.Milliseconds(), merge_comp)
}
