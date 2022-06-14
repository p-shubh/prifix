package main

import (
	"fmt"
	"strings"
)

func main() {

	// var str1 []string

	Input := []string{"pay", "attention", "practice", "attend", "attendence", "pay", "attention", "practice", "attend", "attendence"}

	find := "at"
	// fmt.Println(find)

	fmt.Println("the count is ", call(Input, find))

}

func call(Input []string, find string) int {

	var count, check int

	for i := 0; i < len(Input); i++ {

		find := strings.Split(find, "")

		str1 := strings.Split(Input[i], "")
		// fmt.Println(str1)
		count = 0

		for j := 0; j < len(find); j++ {
			if find[j] == str1[j] {
				count = count + 1
			}
		}

		if count == len(find) {
			check = check + 1
		}

	}

	return check
}
