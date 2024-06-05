package main

import (
	"fmt"
	"strconv"
)

func main(){

	str := "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
	fmt.Println(PrinterError(str))

}

func PrinterError(str string) string {
	 
	fmt.Println()
	strLen := len(str)
	var errorCount int
	for i:=0; i < strLen; i++ {
		if str[i] <= 'm'{
			continue
		}else {
			errorCount++
		}
	}

	return strconv.Itoa(errorCount) + "/"+ strconv.Itoa(strLen)

}