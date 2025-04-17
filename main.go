package main
 
import (
   "fmt"
)

func dividir(dividendo int, divisor int)(int, string){
	if divisor == 0 {
		return 0, "Erro na divisão pro zero"
	} 
    return dividendo / divisor,"Sem erro"
}
func operaçaobasica(a int, b int) (int, int, int){
	soma := a + b
	multiplicaçao := a * b
	subtraçao := a - b
	return soma,multiplicaçao,subtraçao
}
func main(){
	resultado, erro:= dividir(10,2)

	if erro != "sem erro" {
		fmt.Println(erro)
	}else {
		fmt.Println("o resultado da divisão é:", resultado)
	}
	soma,mult,sub := operaçaobasica(10,2)
		fmt.Println(soma)
		fmt.Println(mult)
		fmt.Println(sub)
	
}


   