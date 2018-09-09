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

	f := strings.Split(line, " ") //スペース1文字発見したら分割, fは配列
	n := mustParseInt(f[0])
	a := mustParseInt(f[1])
	b := mustParseInt(f[2])

	var ans []int
	for num := 1; num <= n; num++ {
		num_str := strconv.Itoa(num)
		num_str_array := strings.Split(num_str, "")
		num_int_array := make([]int, len(num_str_array))
		//intのarrayへ
		for i, v := range num_str_array {
			num_int_array[i], _ = strconv.Atoi(v)
		}
		sum := sumArray(num_int_array)
		if sum >= a && sum <= b {
			ans = append(ans, num)
		}
	}
	fmt.Println(sumArray(ans))
}
