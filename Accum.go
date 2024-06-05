package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "ZpglnRxqenU"
	fmt.Println(Accum(str))

}

func Accum(s string) string {
    // your code
	var a strings.Builder
	for i, v := range s{
		x := strings.ToLower(string(v))
		aa := strings.Title(strings.Repeat((x),i+1))
		a.WriteString(aa)
		if i+1 == len(s){
			break
		}
		a.WriteString("-")

	}
	return a.String()
}