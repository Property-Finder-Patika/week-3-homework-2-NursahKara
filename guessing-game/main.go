package main

import (
	"fmt"
	"math/rand"
	"os"
)

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func splitToDigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)
	return ret
}

func main() {
	var input int
	var countSame int
	var countDiff int
	min := 1000
	max := 9999

	random := rand.Intn(max-min) + min

	fmt.Scanf("%d", &input)

	if input == random {
		fmt.Println("YOU ARE WIN")
		os.Exit(0)
	}

	for x, i := range splitToDigits(input) {

		for y, j := range splitToDigits(random) {
			if i == j {
				if x == y {
					countSame++
				} else {
					countDiff++
				}
			}
		}

	}
	if countSame == 4 && countDiff == 0 {
		fmt.Println("YOU ARE WIN")
	} else {
		fmt.Println("Hint: ", "+", countSame, "-", countDiff)
	}

}
