package main

import (
	"fmt"
	"io"
	"os"
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
    contents := make([][]byte, 0)

    for _, line := range strings.Split(fileContents, "\n") {
        if len(line) == 0 {
            continue
        }
        
        contents = append(contents, []byte(line))
    }

    fmt.Println(partOne(contents))
}

// 18:55 - 20:19 | 01 hour 24 minutes
func partOne(contents [][]byte) int {
    reference := []byte("XMAS")

    count := 0

    for i := 0; i < len(contents); i++ {
        for j := 0; j < len(contents[i]); j++ {
            if contents[i][j] != reference[0] {
                continue
            }

            checkTopDown(contents, i, j, reference, &count)
            checkLeftRight(contents, i, j, reference, &count)
            checkDownTop(contents, i, j, reference, &count)
            checkRightLeft(contents, i, j, reference, &count)

            checkTopLeft(contents, i, j, reference, &count)
            checkDownLeft(contents, i, j, reference, &count)
            checkTopRight(contents, i, j, reference, &count)
            checkDownRight(contents, i, j, reference, &count)

            fmt.Printf("%d,%d\n", i, j)
            fmt.Println(count)
        }
    }

    return count
}

func checkLeftRight(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for j := jStart; j < jStart + len(reference) && j < len(contents[iStart]); j++ {
        fmt.Printf("%c", contents[iStart][j])

        if contents[iStart][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        refIndex++
    }

    fmt.Printf("\n")

    if foundXmas {
        (*count)++
    }
}


func checkRightLeft(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for j := jStart; j > jStart - len(reference) && j >= 0; j-- {
        fmt.Printf("%c", contents[iStart][j])
        if contents[iStart][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}


func checkTopDown(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i := iStart; i < iStart + len(reference) && i < len(contents); i++ {
        fmt.Printf("%c", contents[i][jStart])
        if contents[i][jStart] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }


        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

func checkDownTop(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i := iStart; i > iStart - len(reference) && i >= 0; i-- {
        fmt.Printf("%c", contents[i][jStart])
        if contents[i][jStart] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

func checkDownRight(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i, j := iStart, jStart; i < iStart + len(reference) && j < jStart + len(reference) && i < len(contents) && j < len(contents[i]); i += 1 {
        fmt.Printf("%c", contents[i][j])
        if contents[i][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        j++
        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

func checkTopLeft(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i, j := iStart, jStart; i > iStart - len(reference) && j > jStart - len(reference) && i >= 0 && j >= 0; i-- {
        fmt.Printf("%c", contents[i][j])
        if contents[i][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        j--
        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

func checkTopRight(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i, j := iStart, jStart; i > iStart - len(reference) && j < jStart + len(reference) && i >= 0 && j < len(contents[i]); i-- {
        fmt.Printf("%c", contents[i][j])
        if contents[i][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        j++
        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

func checkDownLeft(contents [][]byte, iStart int, jStart int, reference []byte, count *int) {
    refIndex := 0

    foundXmas := false

    for i, j := iStart, jStart; i < iStart + len(reference) && j >= jStart - len(reference) && i < len(contents) && j >= 0; i++ {
        fmt.Printf("%c", contents[i][j])
        if contents[i][j] != reference[refIndex] {
            break
        } else if refIndex == len(reference) - 1 {
            foundXmas = true
        }

        j--
        refIndex++
    }

    fmt.Printf("\n")
    if foundXmas {
        (*count)++
    }
}

