package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	a := mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()
	b := mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()
	c := mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()
	x := mustParseInt(line)

	count := 0
	for ai := 0; ai <= a; ai++ {
		for bi := 0; bi <= b; bi++ {
			for ci := 0; ci <= c; ci++ {
				if ai*500+bi*100+ci*50 == x {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
