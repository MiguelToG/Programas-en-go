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

archivo,err := os.Open("./holaa.txt")
defer func () {
	archivo.Close()
	fmt.Println("Defer")
	r := recover()
	fmt.Println(r)
}()

if err != nil{
panic(err)
}

scanner	:= bufio.NewScanner(archivo)
	
for scanner.Scan(){
	linea := scanner.Text()
	fmt.Println(linea)
}

return true
}