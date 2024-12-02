package main

import (
	"fmt"
	"io"
	"os"
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

    fileContents := string(b)
    reports := make([][]int, 0)

    for _, line := range strings.Split(fileContents, "\r\n") {
        if len(line) == 0 {
            continue
        }
        
        levels := []int{}
        for _, level := range strings.Split(line, " ") {
            num, err := strconv.Atoi(level)
            if err != nil {
                panic(err)
            }

            levels = append(levels, num)
        }

        reports = append(reports, levels)
    }

    count := partOne(reports)
    fmt.Println(count)
    count = partTwo(reports)
    fmt.Println(count)
}

// 20:55 - 21:48 | 53 minutes
func partOne(reports [][]int) int {
    count := 0

    for _, report := range reports {
        isIncreasing := false

        if report[0] - report[len(report)-1] < 0 {
            isIncreasing = true
        }

        isSafe := true

        for i := 0; i < len(report) - 1; i++ {
            if isIncreasing {
                if report[i+1] - report[i] < 1 || report[i+1] - report[i] > 3 {
                    isSafe = false
                    break
                }
            } else {
                if report[i] - report[i+1] < 1 || report[i] - report[i+1] > 3 {
                    isSafe = false
                    break
                }
            }
        }

        if isSafe {
            count++
        }
    }

    return count
}

// 21:51 - 22:04 | 13 minutes
func partTwo(reports [][]int) int {
    count := 0

    for _, report := range reports {
        isSafe := false

        if processReportPartTwo(report) {
            count++
            continue
        }

        for i, _ := range report {
            strippedReport := removeElement(report, i)

            if processReportPartTwo(strippedReport) {
                isSafe = true
                break
            }
        }

        if isSafe {
            count++
        }
    }

    return count
}

func removeElement(report []int, index int) []int {
    strippedReport := []int{}

    for i, v := range report {
        if i == index {
            continue
        }

        strippedReport = append(strippedReport, v)
    }

    return strippedReport
}

func processReportPartTwo(report []int) bool {
    isIncreasing := false

    if report[0] - report[len(report)-1] < 0 {
        isIncreasing = true
    }

    isSafe := true

    for i := 0; i < len(report) - 1; i++ {
        if isIncreasing {
            if report[i+1] - report[i] < 1 || report[i+1] - report[i] > 3 {
                isSafe = false
                break
            }
        } else {
            if report[i] - report[i+1] < 1 || report[i] - report[i+1] > 3 {
                isSafe = false
                break
            }
        }
    }

    return isSafe
}
