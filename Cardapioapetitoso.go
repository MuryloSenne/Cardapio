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
      fmt.println("")
  }
}
