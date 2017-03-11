package main 

import(
"net/http"
"encoding/json"
)
type curso struct{
	Title string `json:"titulo"`
 	NumeroDeVideos int `json:"numero_videos"`
} 
type cursos []curso

func main() {
http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
cursos := cursos{
curso	{"Curso de go",30},
curso	{"Curso de ruby",30},
curso	{"Curso de nodejs",30},
}
json.NewEncoder(w).Encode(cursos)
}