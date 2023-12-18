package main

import "fmt"


func main(){
  var prato string
  fmt.println("Digite seu prato...")
  fmt.println("P - Pizza")
  fmt.println("H - Hamburguer")
  fmt.println("B - Bife com fritas")
  fmt.println("S- Salada Caesar")
  fmt.println("F- Salda de frutas")
  fmt.println("E - Estrogonofe")
  fmt.println("O - Outros")
  fmt.Scan(&prato)

  switch prato{
    case "B":
      fmt.println("Com batata ou maionese ?")
    case "P":
      fmt.println("Calabresa OU Mussarela ?")
    case "H":
      fmt.println("Com queijo ou ovo ?")
    case "S":
      fmt.println("Alface ou rucula ?")
    case "F":
      fmt.println("Banana ou Maça")
    case "E":
      fmt.println("Carne ou queijo")
    case "O":
      fmt.println("Não gostou do nosso cardapio")    
    default:
    case "H":
    fmt.println("Não entendi seu paradar !! ")
  }
}
