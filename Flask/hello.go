package main
import "fmt"
func main(){
  arreglo:=[7]int{0,1,4,6,10,9}
  for i, j:= range arreglo{
    fmt.Printf("Valor de j: %d en vuelta #%d\n", j,i)
  }
  for i:= range arreglo{
    fmt.Printf("Valor de i: %d\n", i)
  }
}