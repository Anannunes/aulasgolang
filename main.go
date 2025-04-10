package main

import (
"fmt"
"sort"
)

func main() {
	var ages = []int{17, 16, 20, 40}
	names := []string{"mario", "luigi", "homem aranha", "batman"}
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "caroline"))

	for x:=0; x > 5; x++ {
		fmt.Println(x)
	}
	for i := 1; i < len(names); i++ {
		fmt.Println(names[i])

	}
		
	for index, value := range ages {
			fmt.Println(" o indice é", index," e o valor", value)
	}
}
