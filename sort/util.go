package sort

import "fmt"

func Swap(data []int, i, j int) {
	var tmp int

	tmp = data[i]
	data[i] = data[j]
	data[j] = tmp
}

func Print(data []int, name string) {
	fmt.Printf("%s: ", name)

	for _, i := range data {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}