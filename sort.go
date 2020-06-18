package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 7, 5, 4, 3, 2}
	letters := []string{"c", "g", "t", "i"}
	sort.Ints(nums)
	fmt.Println(nums)
	sort.Strings(letters)
	fmt.Println(nums)
	fmt.Println(letters)

}
