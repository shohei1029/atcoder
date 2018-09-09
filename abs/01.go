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
	num, err := strconv.Atoi(text) //文字以外が入ってきたらエラー?
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

	f := strings.Split(line, " ") //スペース1文字発見したら分割, fは配列
	b := mustParseInt(f[0])
	c := mustParseInt(f[1])

	scanner.Scan()
	text := scanner.Text()

	fmt.Println(a+b+c, text)
}
