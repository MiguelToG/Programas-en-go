package main

import "fmt"
type user struct{
		edad int
		nombre,apelldo string
		
	}
	func (usuario user)nombre_completo()string {
		return usuario.nombre +" "+usuario.apelldo
	}
	func (this *user) set_name(n string){
    this.nombre = n


	}
func main() {
	var miguel user
miguel.nombre="miguel"
miguel.set_name("mono")
	fmt.Println(miguel.nombre)
}