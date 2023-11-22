package main

// http
import (
	"fmt"

	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello k8s")

}

func main() {

	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":8888", nil)

}
