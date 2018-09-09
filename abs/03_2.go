package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mustParseInt(text string) int {
	num, err := strconv.Atoi(text) //数字としてパースできなかったらエラー
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string
	scanner.Scan()
	line = scanner.Text()

	f := strings.Split(line, "")
	a := mustParseInt(f[0])
	b := mustParseInt(f[1])
	c := mustParseInt(f[2])

	fmt.Println(a + b + c)
}
