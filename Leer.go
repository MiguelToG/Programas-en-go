package main 
import (
"io/ioutil"
"fmt"

)
func main() {
dile_data,err :=	ioutil.ReadFile("C:\\Users\\Miguel\\Documents\\GitHub\\Programas-en-go\\hola.txt")
if (err != nil){
fmt.Println("Hubo un error")
}
fmt.Println(string(dile_data))
}