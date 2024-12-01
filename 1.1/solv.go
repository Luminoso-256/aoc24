package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lList []int
	var rList []int

	b, _ := os.ReadFile("input.txt")
	for _, ln := range strings.Split(string(b), "\n") {
		p := strings.Fields(ln)
		l, _ := strconv.Atoi(p[0])
		fmt.Printf("%v\n", p[1])
		r, _ := strconv.Atoi(strings.Trim(p[1], " "))
		lList = append(lList, l)
		rList = append(rList, r)
	}
	sort.Ints(lList)
	sort.Ints(rList)
	sum := 0
	for i, e := range lList {
		fmt.Printf("%d - %d\n", e, rList[i])
		sum += int(math.Abs(float64(e - rList[i])))
	}
	fmt.Printf("%d\n", sum)
}
