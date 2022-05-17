package sayhello

/*
hello.go for printing the hello world message on browser from server.
*/
import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type InputString struct {
	HelloWorld string
}
type SayHello interface {
	Hello() InputString
}

/*
Hello () for processing the request from user and sending the "hello world" response to user.
*/

func (n *InputString) Hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	result := strings.TrimSpace(strings.ToUpper(vars["text"]))
	// comparison between user input and expected output.
	if result == "HELLO WORLD" {
		var output InputString = InputString{HelloWorld: result}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(output.HelloWorld)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Wrong Input!!!", 400)
	}

}
