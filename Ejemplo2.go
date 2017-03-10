package main
import ("fmt"
"bufio"
"os")
 func main() {
 	c := true
	fmt.Printf("Hay fay %v \n",c)
	a := 14.2
	fmt.Printf("Ho lo %f \n",a)
	fmt.Printf("mi nombre es %s","Miguel \n")
	var edad int
	fmt.Println("Ingresa tu edad")
    fmt.Scanf("%d\n",&edad)
    fmt.Printf("Mi edad es %d\n",edad) 
    var nombre string 
    fmt.Println("Coloca tu nombre:")
    fmt.Scanf("%v\n",&nombre)
    fmt.Printf("Mi nombre es %s\n",nombre)
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("ingresa tu nombre")
    text,err := reader.ReadString('\n')
    if err != nil {
    	fmt.Println(err)
    }else{
    	fmt.Println("Hola"+text)
    }
    
}