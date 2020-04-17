package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val1 = "20"
	var val2 = "10"
	var strval1, err1 = strconv.Atoi(val1)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	var strval2, err2 = strconv.Atoi(val2)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(strval1 + strval2)

}
