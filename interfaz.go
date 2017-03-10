package main 
import "fmt"
type User interface{
permisos() int
Nombre() string
}
type admin struct{
nombre string
}
func (this admin) permisos()int{
return 5
}
func (this admin) Nombre()string{
return this.nombre
}
type Editor struct{
nombre string
}
func (this Editor)permisos()int{
return 3
}
func (this Editor) Nombre()string{
return this.nombre
}

func auth(user User) string{
	permisos := user.permisos()
if permisos>4 {
	return user.Nombre()+" Tiene permisos de administrador"
} else if permisos == 3{
	return user.Nombre()+" Tiene permisos de editor"
}
return ""
}



func main() {
/*	admin := admin{"Miguel"}
	editor := Editor{"Cuco"}
	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
 */
usuarios := []User{admin{"Negro"},Editor{"No"}}
for _,usuario := range usuarios{

	fmt.Println(auth(usuario))
}


}