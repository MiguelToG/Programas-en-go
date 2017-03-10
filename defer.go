package main 
import (
"fmt"
"bufio"
"os"
)

func main() {
ejecucion :=	readFile()
fmt.Println(ejecucion)
}
func readFile() bool{

archivo,err := os.Open("./hola.txt")
defer func () {
	archivo.Close()
	fmt.Println("Defer")
}()
if err != nil{
	fmt.Println("Hubo error")
}

scanner	:= bufio.NewScanner(archivo)
	
for scanner.Scan(){
	linea := scanner.Text()
	fmt.Println(linea)
}

return true
}