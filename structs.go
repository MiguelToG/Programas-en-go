package main

import "fmt"
type user struct{
		edad int
		nombre,apelldo string
		
	}
func main() {
	var miguel user

	fmt.Println(miguel)

fmt.Println(user{nombre : "Negro",apelldo : "mornogo"})
negru := user{nombre : "Negro",apelldo : "mornogo"}
fmt.Println(negru)
tomi := user{0,"",""}
fmt.Println(tomi)
pico := new(user)
fmt.Println(pico)
}