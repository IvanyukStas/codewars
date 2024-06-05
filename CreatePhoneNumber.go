package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main(){ 	
	arr := [10]uint{1,2,3,4,8,6,7,1,9,0}
	fmt.Println(CreatePhoneNumber(arr))
}

func CreatePhoneNumber(arr [10]uint) string  {
	var str string
	for _, v := range arr {
		str += strconv.Itoa(int(v))
	}
	
	re, _ := regexp.Compile(`(?P<val1>\d{3})(?P<val2>\d{3})(?P<val3>\d{4})`)
template := "(${val1}) ${val2}-${val3}"
fmt.Println(re.ReplaceAllString(str, template))
return re.ReplaceAllString(str, template)
}