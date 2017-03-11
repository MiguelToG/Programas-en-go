package main
import (
"fmt"
"bufio"
"os"
"strings"
)
func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("			*****Voltear cadenas****")
	fmt.Println(">>>Ingresa la cadena a voltear")
text,err := scanner.ReadString('\n')

if (err != nil){
panic (err)
}

var letra [50]string
i := 0
letras := strings.Split(text,"")
for _,letraso := range (letras){
		letra[i]=letraso
		i++
	}
	fmt.Println(">>>La cadena al revés es :")
	
for x := i; x >= 0; x-- {
fmt.Print(letra[x])	
}

}