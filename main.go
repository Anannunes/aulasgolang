package main
	
import "fmt"   

func main(){ 
	var senha int
	var usuario string
	fmt.Println("Digite seu nome de usuario:")
	fmt.Scan(&usuario)
	fmt.Println("Digite sua senha:")
	fmt.Scan(&senha)
	if usuario == "admin" && senha==1234 {
	fmt.Println("pode entrar")
} else {
    fmt.Println("acesso negado,tentar novamente")
}
}