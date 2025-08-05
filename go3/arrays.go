package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20

	numaStr := "12345"
	num, err := strconv.Atoi(numaStr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println(" parsed integer:", num)
	fmt.Println(" parsed integer:", num+1)

}
