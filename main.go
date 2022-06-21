package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Print(numbers int) string {
	var res []string
	for i := 1; i <= numbers; i++ {
		if i%3 == 0 {
			if i%5 == 0 {

				res = append(res, "Frontend Backend")

				continue
			}
			res = append(res, "Frontend")
			continue

		} else if i%5 == 0 {
			res = append(res, "Backend")
			continue
		}

		res = append(res, strconv.Itoa(i))
		continue
	}
	return strings.Join(res, ", ")

}

func main() {
	fmt.Println(Print(50))
}
