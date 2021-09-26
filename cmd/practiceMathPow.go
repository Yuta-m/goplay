package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var n1, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	var n2, _ = strconv.Atoi(sc.Text())

	fmt.Print(math.Sqrt(math.Pow(float64(n1-n2), 2.0)))
}
