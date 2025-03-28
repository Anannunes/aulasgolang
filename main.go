package main
	
import "fmt"   

func main(){ 
	var a int
	var b int
	fmt.Println("Digite um numero:")
	fmt.Scan(&a)
	fmt.Println("digite outro numero:")
	fmt.Scan(&b)
	fmt.Println(" A soma é: ", a + b)
	fmt.Println("A subtração é:", a - b)
	fmt.Println("a divisão é:", a / b)
	fmt.Println("A multiplicação é:", a * b)
	fmt.Println("O resto da divisão: ", a % b)

}