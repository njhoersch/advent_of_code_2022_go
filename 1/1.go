package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partTwo()
}

func findElfWithMostCals() {
	file, err := os.Open("1input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cals := []int{}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			cals = append(cals, total)
			total = 0
			continue
		}
		total = total + num
	}

	if scanner.Err() != nil {
		fmt.Println("Error scanning files: ", err)
		return
	}

	max := cals[0]
	index := 0

	for i, val := range cals {
		if val > max {
			max = val
			index = i
		}
	}

	fmt.Printf("Elf %d has the highest calories.\n", index+1)
	fmt.Printf("Elf %d has %d calories.\n", index+1, max)
}

// Define a type for the heap elements
type MaxHeap []int

// Implement the heap.Interface methods for MaxHeap
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Use ">" for a max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func partTwo() {
	file, err := os.Open("1input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// cals := []int{}
	h := &MaxHeap{}
	heap.Init(h)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			heap.Push(h, total)
			total = 0
			continue
		}
		total = total + num
	}

	if scanner.Err() != nil {
		fmt.Println("Error scanning files: ", err)
		return
	}

	val := 0

	for i := 0; i < 3; i++ {
		val = val + heap.Pop(h).(int)
	}

	fmt.Println(val)
}
