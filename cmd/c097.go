package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	var s = strings.Split(sc.Text(), " ")
	var n, _ = strconv.Atoi(s[0])
	var x, _ = strconv.Atoi(s[1])
	var y, _ = strconv.Atoi(s[2])

	for i := 1; i <= n; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println("AB")
		} else if i%x == 0 {
			fmt.Println("A")
		} else if i%y == 0 {
			fmt.Println("B")
		} else {
			fmt.Println("N")
		}
	}
}
