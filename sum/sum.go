package sum

import (

)

func sum(array []int) int {
	var holder int
	for _, val := range array {
		holder += val
	}
	return holder
}

func sumAll(arrays ...[]int) []int {
	var result []int
	for _, val := range arrays {
		result = append(result, sum(val))
	}
	return result
}

func sumAllTails(arrays ...[]int) []int {
	var result []int
	for _, val := range arrays {
		if len(val) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, sum(val[1:]))
		}
	}
	return result
}