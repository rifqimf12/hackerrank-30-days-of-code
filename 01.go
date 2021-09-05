package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// mycode{
	var (
		ii uint64
		dd float64
		ss string
	)
	scanner.Scan()
	ii, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	dd, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	ss = scanner.Text()
	fmt.Printf("%d\n%.1f\n%s\n", i+ii, d+dd, s+ss)
	// }mycode
}
