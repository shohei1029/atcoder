package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	b := make([]byte, 1000000)
	r.Buffer(b, 1000000)
	r.Scan()
	var s string
	s = r.Text()

	t := ""

	prev_s := ""

	for {

		s = strings.TrimSuffix(s, "dreamer")
		s = strings.TrimSuffix(s, "dream")
		s = strings.TrimSuffix(s, "eraser")
		s = strings.TrimSuffix(s, "erase")
		if s == t {
			fmt.Println("YES")
			break
		} else if s == prev_s {
			fmt.Println("NO")
			break
		}

		prev_s = s
	}
}
