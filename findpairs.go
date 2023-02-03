package main

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
