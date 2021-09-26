package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var s = strings.Split(sc.Text(), " ")
	var people, _ = strconv.Atoi(s[0])
	var length, _ = strconv.Atoi(s[1])
	var result int
	scores := []int{}
	intervals := []int{}

	for i := 0; i <= people; i++ {
		score := 100
		for j := 0; j < length; j++ {
			sc.Scan()
			var interval, _ = strconv.Atoi(sc.Text())
			if i == 0 {
				intervals = append(intervals, interval)
				continue
			}
			tolerance := math.Sqrt(math.Pow(float64(interval-intervals[j]), 2.0))
			if tolerance == 0 {
				// do nothing
			} else if tolerance <= 10 {
				score -= 1
			} else if tolerance <= 20 {
				score -= 2
			} else if tolerance <= 30 {
				score -= 3
			} else {
				score -= 5
			}
		}
		if i != 0 {
			if score < 0 {
				score = 0
			}
			scores = append(scores, score)
		}
	}
	for k := 0; k < people; k++ {
		score := scores[k]
		if result < score {
			result = score
		}
	}
	fmt.Println(result)
}
