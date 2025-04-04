package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){ 
	greeting :="Hello my friends!"
	fmt.Println(strings.Contains(greeting,"dog"))
    fmt.Println(strings.ReplaceAll(greeting,"my", "mine"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"friends"))
	ages := []int{50, 80, 10}
    sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 80)
	fmt.Println(index)	
	names := []string{"caroline","maicon","diego"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names,"caroline"))

}
