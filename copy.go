package main 
import "fmt"
func main() {
	/*
     Copia el minimo de elementos 

	*/
slice := []int{1,2,3,4}
copia := make([]int,len(slice),cap(slice)*2)//Para evitar errores, pasa el tamaño del slice anterior
copy(copia,slice)

fmt.Println(slice)
fmt.Println(copia)
}
