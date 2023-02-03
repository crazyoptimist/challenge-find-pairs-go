package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindPairs(nums []int, target int) (result [][]int) {
	hashmap := make(map[int]int)
	for idx, num := range nums {
		if _, found := hashmap[target-num]; found {
			result = append(result, []int{num, target - num})
		}

		hashmap[num] = idx
	}
	return
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: <app> <integers comma separated> <target integer>")
		return
	}

	nums := make([]int, 0)
	target, _ := strconv.Atoi(os.Args[2])

	for _, numberString := range strings.Split(os.Args[1], ",") {
		num, _ := strconv.Atoi(numberString)
		nums = append(nums, num)
	}

	fmt.Println(FindPairs(nums, target))

}
