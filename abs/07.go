package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string
	scanner.Scan()
	line = scanner.Text()

	//	n := mustParseInt(line)

	scanner.Scan()
	line = scanner.Text()

	cards_str_array := strings.Split(line, " ")
	cards_int_array := make([]int, len(cards_str_array))
	//intのarrayへ
	for i, v := range cards_str_array {
		cards_int_array[i], _ = strconv.Atoi(v)
	}
	cards_slc := cards_int_array[:] //絶対これ冗長....もとからsliceじゃね...

	sort.Ints(cards_slc)

	alice_score := 0
	bob_score := 0

	for { //気持ち悪い..
		alice_score += cards_slc[len(cards_slc)-1]
		cards_slc = cards_slc[:len(cards_slc)-1] //末尾削除
		if len(cards_slc) == 0 {
			break
		}
		bob_score += cards_slc[len(cards_slc)-1]
		cards_slc = cards_slc[:len(cards_slc)-1] //末尾削除
		if len(cards_slc) == 0 {
			break
		}

	}

	fmt.Println(alice_score - bob_score)
}
