package main

import "fmt"
type human struct{
name string

}
func (this human) hablar()string{
return "bla bla bla"
}
type tutor struct{
human
}
func (this tutor)hablar()string{
	return this.human.hablar()+"Soy Miguel"
}
func main() {
	tutor := tutor{human{"Miguel"}}
fmt.Println(tutor.hablar())
}