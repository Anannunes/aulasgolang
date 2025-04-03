package main
	
import "fmt"   

func main(){ 
	var ages = [4]int{17,16,20,40}
	nomes := [4]string{"mario","luigi","homem aranha","batman"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[3]="Miles Morales"
	fmt.Println(nomes)
}
