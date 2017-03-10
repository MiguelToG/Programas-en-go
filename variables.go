package main 

import ("fmt"
"strconv"

)
func main() {
i := 22
i_str := strconv.Itoa(i)
fmt.Println("hola"+i_str)
f := "22"
f_int,_ := strconv.Atoi(f)
fmt.Println(f_int+1) 	
}