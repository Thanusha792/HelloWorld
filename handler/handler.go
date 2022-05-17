package handler

/*
 handler.go for handling the requests.
*/
import (
	"HelloWorld/configuration"
	sayhello "HelloWorld/sayHello"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Handlerfunc () for handling the endpoints based on user request.
*/
func Handlerfunc() {

	fmt.Println("Welcome")
	r := mux.NewRouter()
	var inputString sayhello.InputString
	r.HandleFunc("/getText/{text}", inputString.Hello).Methods("GET")
	port := configuration.PortCoonect()
	log.Fatal(http.ListenAndServe(port.PortNumber, r))
}
