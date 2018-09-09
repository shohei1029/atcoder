package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//エラーが発生したときは勝手にpanicしろって関数にはmustって名前を付ける慣例
func mustParseInt(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func uniq(a []int) []int {
	m := make(map[int]bool)
	uniq_a := []int{}

	for _, elem := range a {
		if !m[elem] {
			m[elem] = true
			uniq_a = append(uniq_a, elem)
		}
	}
	return uniq_a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string
	scanner.Scan()
	line = scanner.Text()

	n := mustParseInt(line)

	mochis := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		mochis[i] = mustParseInt(scanner.Text())
	}

	fmt.Println(len(uniq(mochis)))
}
