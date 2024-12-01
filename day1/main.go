package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("ERROR READING FILE")
		return
	}

	b, err := io.ReadAll(file)

	membersLeft := []int{}
	membersRight := []int{}

	lines := strings.Split(string(b), "\r\n")

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		if len(parts) < 2 {
			continue
		}

		left, _ := strconv.Atoi(parts[0])
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		membersLeft = append(membersLeft, left)
		membersRight = append(membersRight, right)
	}

	slices.Sort(membersLeft)
	slices.Sort(membersRight)

	sum := 0
	for _, v := range membersLeft {
		sum += v * findCount(membersRight, v)
	}

	fmt.Println(sum)
}

func findCount(slice []int, val int) int {
	count := 0
	for _, val2 := range slice {
		if val == val2 {
			count++
		}
	}

	return count
}
