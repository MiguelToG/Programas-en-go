package main 
import (
"fmt"
"time"
"strings"
)

func main() {
	go mi_nombre_lento("Miguel")
	fmt.Println("Esperando")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lento(name string){
	letras := strings.Split(name,"")
	for _,letra := range (letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}


}