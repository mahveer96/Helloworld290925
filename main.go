package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


// ✅ Define the Response struct for JSON output
type Response struct {
	Message string `json:"message"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {


	// Dummy hardcoded credentials
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "Login successful"})

}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world welcome to Golang")
}

func main(){
	fmt.Println("Hello world welcome to Golang ")
    http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))


}