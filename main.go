package main
 
import (
   "fmt")

func main(){
   capitais ;=map[string]string{ 
   "SP" : "São Paulo"
   "RJ" : "Rio de Janeiro "
   "ES" : "Espirito Santo"
}
 capitais["BH"] = "Belo Honizonte"

 for k,v := range capitais{
   fmt.Println("Sigla, Nome", k, v)
 }
 delete(capitais, "AC")

 fork, v :=range capitais{
   fmt.Println?("Sigla, Nome", k, v)
   {

   }
 }