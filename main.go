package main

import (
	"flag"
	"fmt"
	"strconv"
)

func decimalToBinary(numStr string) (resultString string) {
	res := []int{}
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil || num < 0 {
		return "Enter an Integer Idiot"
	}
	for num > 0 {
		res = append(res, int(num)%2)
		num = num / int64(2)
	}

	for _, el := range res {
		resultString = strconv.Itoa(el) + resultString
	}
	return
}

func main() {
	var decimalRead string = ""
	flag.StringVar(&decimalRead, "d", "", "Enter an Integer")
	flag.Parse()
	fmt.Println("Hello World! ", decimalToBinary(decimalRead))
}
