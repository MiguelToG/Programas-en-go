package main 
import "fmt"
func main() {
	y := 10
	c :=12
	if y<c {
		fmt.Printf("%d Es mayor a %d",c,y)
	}else if y>c {
		fmt.Println("Entré al else if")
	}else{
		fmt.Println("Entré al else")
	}
   
}