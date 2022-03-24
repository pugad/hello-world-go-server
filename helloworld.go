package main
import (
	"fmt"
	"net/http"
)

func main() {
	
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world! You've requested: %s\n", r.URL.Path)
    })

	fmt.Println("Serving at 0.0.0.0:80")

	http.ListenAndServe("0.0.0.0:80", nil)
}