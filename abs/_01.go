package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mustParseInt(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	var line string
	scanner.Scan()
	line = scanner.Text()
	sum += mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()
	nums := strings.Split(line, " ")
	sum += mustParseInt(nums[0]) + mustParseInt(nums[1])

	scanner.Scan()
	line = scanner.Text()

	fmt.Println(sum, line)
}
