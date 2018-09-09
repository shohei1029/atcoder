package main

import (
	"bufio"
	"fmt"
	"os"
)

func myDaydream(s, t string) {
	if t == s {
		fmt.Println("YES")
		os.Exit(0)
	} else if len(t) > len(s) {
		return
	} else {
		sippo := []string{"dream", "dreamer", "erase", "eraser"}
		for _, v := range sippo {
			myDaydream(s, t+v)
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	b, _ := r.ReadBytes('\n')
	s := string(b[:len(b)-1])

	t := ""

	myDaydream(s, t)
	fmt.Println("NO")
}
