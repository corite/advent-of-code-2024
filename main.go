package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func d1p1() {
	file, err := os.ReadFile("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	diff := make([]int, len(lines))

	for index, line := range lines {
		parts := strings.Split(line, "   ")
		left[index], _ = strconv.Atoi(parts[0])
		right[index], _ = strconv.Atoi(parts[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for index := range len(lines) {
		if left[index] < right[index] {
			diff[index] = right[index] - left[index]
		} else {
			diff[index] = left[index] - right[index]
		}
		sum = sum + diff[index]
	}
	fmt.Printf("total sum of diffs: %d\n", sum)
}
func d1p2() {
	file, err := os.ReadFile("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	assoc := make(map[int]int)
	for index, line := range lines {
		parts := strings.Split(line, "   ")
		left[index], _ = strconv.Atoi(parts[0])
		right[index], _ = strconv.Atoi(parts[1])
	}
	for _, item := range right {
		assoc[item] = assoc[item] + 1
	}
	total := 0
	for _, item := range left {
		total = total + assoc[item]*item
	}
	fmt.Printf("similarity score: %d\n", total)
}

func main() {
	d1p1()
	d1p2()
}
