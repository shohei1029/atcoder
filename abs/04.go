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

	n := mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()

	f := strings.Split(line, " ") //スペース1文字発見したら分割, fは配列

	//intのリストを作る
	ary := make([]int, n)
	for i := range ary {
		ary[i], _ = strconv.Atoi(f[i])
	}

	count := 0
	for {
		//		for i := 0; i < len(ary); i++ {
		for i := range ary {
			if ary[i]%2 == 0 {
				ary[i] /= 2
			} else {
				fmt.Println(count)
				os.Exit(0)
			}
		}
		count += 1
	}
}
