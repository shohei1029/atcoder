package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//エラーが発生したときは勝手にpanicしろって関数にはmustって名前を付ける慣例
func mustParseInt(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string
	scanner.Scan()
	n := mustParseInt(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		f := strings.Split(line, " ")

		t := mustParseInt(f[0])
		x := mustParseInt(f[1])
		y := mustParseInt(f[2])

		if x+y > t {
			fmt.Println("No")
			return
		} else if t%2 != (x+y)%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
