package main
import (
	"vania_16119464_pert4/handler" //sesuaikan dengan nama folder (case sensitive)
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/api/", handler.API)
	//Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
	log.Println("localhost : 8064")
	http.ListenAndServe(":8064", nil) 
}