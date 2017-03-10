package main 
import (
"os"
"fmt"
"bufio"
)
func main() {
archivo,err := os.Open("./hola.txt")
if err != nil{
	fmt.Println("Hubo error")
}
scanner	:= bufio.NewScanner(archivo)
for scanner.Scan(){
	linea := scanner.Text()
	fmt.Println(linea)
}
}