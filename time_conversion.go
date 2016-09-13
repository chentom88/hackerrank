package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)

	time := input[:8]
	meridien := strings.ToUpper(input[8:len(input)])

	timeComps := strings.Split(time, ":")

	hour, _ := strconv.Atoi(timeComps[0])
	if strings.ToUpper(meridien) == "PM" {
		if hour != 12 {
			hour = (hour + 12)
		}
	} else {
		hour = hour % 12
	}

	fmt.Printf("%02d:%s:%s\n", hour, timeComps[1], timeComps[2])
}
