/*
	Written by Eoghan O'Connor
	Eliza chatbot in golang
	3rd year software project 
*/


package main

import(

	"net/http"

)

//serving the webpage
func requestHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "index.html")

}

func main(){

	http.HandleFunc("/",requestHandler)

	http.ListenAndServe(":8080",nil)
}