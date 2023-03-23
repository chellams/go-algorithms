package main

import (
	"sort"
)

type temp struct {
	num  int
	name string
}

func main() {
	var temps []temp

	temps = append(temps, temp{num: 1, name: "A"})
	temps = append(temps, temp{num: 3, name: "B"})
	temps = append(temps, temp{num: 2, name: "C"})
	temps = append(temps, temp{num: 5, name: "D"})
	temps = append(temps, temp{num: 4, name: "E"})
	temps = append(temps, temp{num: 5, name: "F"})

	max := 0
	sort.Slice(temps, func(i, j int) bool {
		if max < temps[i].num {
			max = temps[i].num
		}
		return temps[i].num < temps[j].num
	})

}
