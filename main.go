package main
 
 import "fmt"
 
 func main() {
    var idade int
	fmt.Println("digite dua idade")
	fmt.Scan(&idade)
	if idade <18 {
	fmt.Println("vc é de menor")
	}else if idade >=18 && idade <=60{
		fmt.Println("vc é de maior")
	}else {
		fmt.Println("vc é um idoso")
	}
}