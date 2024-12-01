package main

import (
	"fmt"
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
	simScore := 0
	for _, e := range lList {
		//this is slow for sure but w/e
		count := 0
		for _, r := range rList {
			if r == e {
				count++
			}
		}
		simScore += count * e
	}
	fmt.Printf("%v\n", simScore)
}
