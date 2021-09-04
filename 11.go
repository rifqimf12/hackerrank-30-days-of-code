package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	// my code{
	var arrWidth, arrHeight = len(arr[0]), len(arr)
	const hgHeight, hgWidth, smallestValue = 3, 3, -9
	sumOfHourGlass := func(x, y int) int {
		total := 0
		for i := 0; y+i < y+hgWidth; i++ {
			// upper side
			total += int(arr[x][y+i])
			// bottom side
			total += int(arr[x+hgHeight-1][y+i])
		}
		median := (y + y + hgWidth) / 2
		for i := x + 1; i < x+1+hgHeight-2; i++ {
			total += int(arr[i][median])
		}
		return total
	}
	maxTotal := smallestValue * (hgWidth*2 + hgHeight - 2)
	for i := 0; i < arrHeight; i++ {
		for ii := 0; ii < arrWidth; ii++ {
			if ii+hgWidth-1 < arrWidth && i+hgHeight-1 < arrHeight {
				if total := sumOfHourGlass(i, ii); maxTotal <= total {
					maxTotal = total
				}
			}
		}
	}
	printf("%d\n", maxTotal)
	writer.Flush()
	// }my code
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
