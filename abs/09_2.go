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
	line = scanner.Text()

	f := strings.Split(line, " ")
	n := mustParseInt(f[0])
	total := mustParseInt(f[1])

	for ykt := 0; ykt <= n; ykt++ {
		for hgt := 0; hgt <= n-ykt; hgt++ {
			if ykt*10000+hgt*5000+(n-ykt-hgt)*1000 == total {
				fmt.Println(ykt, hgt, n-ykt-hgt)
				os.Exit(0)
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
