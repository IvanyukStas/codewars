package main

import "fmt"

func main() {
	arr := []int{-5, -456465}

	fmt.Println(MinMax(arr))

}

func MinMax(arr []int) [2]int {
	var vMin, vMax int = arr[0], arr[0]

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