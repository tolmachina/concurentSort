package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func testConcurentSort() {
	k := 4
	fmt.Println(concurrentSort([]int{6, 5, 4, 3, 2}, k))
	fmt.Println(concurrentSort([]int{6, 5, 4, 3, 2, 1}, k))
	fmt.Println(concurrentSort([]int{6, 5, 4, 3, 2, 1, 0}, k))
	fmt.Println(concurrentSort([]int{7, 6, 5, 4, 3, 2, 1, 0}, k))
	fmt.Println(concurrentSort([]int{7, 6, 5, 4, 3, 2, 1, 0, -1, -2}, k))
}

func main() {
	input := userInput()
	k := 4
	// testConcurentSort()
	// testInts := []int{5, 1, 4, 3, 7, 2, 9, 8}
	sortedInput := concurrentSort(input, k)
	fmt.Println("Sorted: ", sortedInput)
}

func userInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integers in one line: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	numStr := strings.Split(text, " ")
	if len(numStr) < 7 {
		fmt.Println("Sort it yourself!")
	}
	var nums []int
	for i := range numStr {
		num, err := strconv.Atoi(numStr[i])
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func concurrentSort(arrInts []int, k int) []int {
	fmt.Println("initial: ", arrInts)
	var wg sync.WaitGroup
	chunkMaxSize := partition(len(arrInts), k)
	twoD := make([][]int, k)
	var partsCount int
	for i := 0; i < len(arrInts); i += chunkMaxSize {
		part := makePart(i, chunkMaxSize, arrInts)
		twoD[partsCount] = part
		partsCount++
		wg.Add(1)
		go func() {
			defer wg.Done()
			bubbleSort(part)
		}()
	}
	wg.Wait()
	for j := 0; j < k; j++ {
		fmt.Println("Pre-Sorted quarter: ", twoD[j])
	}
	
	sortedArr := merge(twoD, k, len(arrInts))
	fmt.Println(sortedArr)
	return sortedArr
}

func makePart(i int, chunkMaxSize int, arrInts []int) []int {
	var part []int
	if i+chunkMaxSize < len(arrInts) {
		part = arrInts[i : i+chunkMaxSize]
	} else {
		part = arrInts[i:]
	}
	return part
}

func Swap(nums []int, index int) {
	nums[index], nums[index+1] = nums[index+1], nums[index]
}

func bubbleSort(nums []int) {
	var swapped bool
	for {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func partition(lenght int, k int) int {
	chunkMaxSize := int(math.Ceil(float64(lenght) / float64(k)))
	return chunkMaxSize
}

type Counter struct {
	counts []int
	maxed  []bool
}

func merge(twoD [][]int, k int, size int) []int {
	fmt.Println("Merging", twoD)
	var sortedArr []int
	counter := Counter{counts: make([]int, k), maxed: make([]bool, k)}
	indexForCounter := 0

	for len(sortedArr) < size {
		smallest := math.MaxInt
		for i := 0; i < len(twoD); i++ {
			if counter.counts[i] < len(twoD[i]) {
				if twoD[i][counter.counts[i]] < int(smallest) {
					smallest = twoD[i][counter.counts[i]]
					indexForCounter = i
				}
			}
		}
		sortedArr = append(sortedArr, smallest)
		counter.counts[indexForCounter] += 1
	}
	return sortedArr
}
