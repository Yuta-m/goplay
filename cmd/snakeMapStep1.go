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
	var row, _ = strconv.Atoi(s[0])
	var crdCounter, _ = strconv.Atoi(s[2])
	var crdMap = map[int][]string{}

	// 座標の生成
	for i := 0; i < row; i++ {
		sc.Scan()
		crdMap[i] = strings.Split(sc.Text(), "")
	}

	for k := 0; k < crdCounter; k++ {
		sc.Scan()
		var coordinate = strings.Split(sc.Text(), " ")
		var y, _ = strconv.Atoi(coordinate[0])
		var x, _ = strconv.Atoi(coordinate[1])
		result := crdMap[y]
		fmt.Println(result[x])
	}
}
