package main

import "fmt"

// write a quicksort function
func quicksort(request []int) []int {
	var response, left, right []int
	length := len(request)
	if length <= 1 {
		return request
	}

	center := request[0]

	for i := 1; i < length; i++ {
		if request[i] > center {
			right = append(right, request[i])
		} else {
			left = append(left, request[i])
		}
	}
	right = quicksort(right)
	left = quicksort(left)
	response = append(append(left, center), right...)
	return response
}

func main() {
	origin := []int{1, 2, 0, 22, 4, 3, 12, 9}
	fmt.Println(quicksort(origin))
}
