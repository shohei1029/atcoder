package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
    s := string(b)

	t := ""

	prev_s := ""
	for {
		if s == t {
			fmt.Println("YES")
			break
		} else if s == prev_s {
			fmt.Println("NO")
			break
		}
		s = strings.Replace(s, "dreamer", "", -1)
		s = strings.Replace(s, "dream", "", -1)
		s = strings.Replace(s, "eraser", "", -1)
		s = strings.Replace(s, "erase", "", -1)
		fmt.Println(s)

		prev_s = s
	}

}
