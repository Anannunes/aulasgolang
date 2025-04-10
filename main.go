package main
 
 import "fmt"
 
 func main() {
     var saldo int
     var escolha int
	 var numero int

	 fmt.Print("digite seu saldo: ")
	 fmt.Scan(&saldo)

	 fmt.Println("Opções disponiveis:")
	 fmt.Println("1 - sacar")
	 fmt.Println("2 - depositar")
	 fmt.Println("escolha uma das opções: ")
	 fmt.Scan(&escolha)

	 if escolha == 1{
		fmt.Println("escolha um valor para sacar")
		fmt.Scan(&numero)
	 }
		if  saldo > numero{
			fmt.Println("saldo insuficiente ")
		} else if saldo >= numero {
			saldo = saldo - numero
			fmt.Println("saque realizado com sucesso")
			fmt.Println("seu saldo atual é ", saldo)
			} 
		
	 if escolha == 2 {
		fmt.Println("escolha o valor que você quer depositar: ")
		fmt.Scan(&numero)
	 saldo = saldo + numero
	 fmt.Println("seu saldo atual é ", saldo)
	}

}