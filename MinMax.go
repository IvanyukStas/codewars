package main

import "fmt"

func main() {
	arr := []int{22, 1}

	fmt.Println(MinMax(arr))

}

func MinMax(arr []int) [2]int {
	var vMin, vMax int = 1, 1

	for _, v := range arr{
		if v < vMin {
			vMin = v
		}
		if  v > vMax{
			vMax = v
		}
	}

	return [2]int{vMin,vMax}
}