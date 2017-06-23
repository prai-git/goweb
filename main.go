// routes.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prai-git/upload"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload.Upload).Methods("GET")
	r.HandleFunc("/fileUpload", upload.UploadFile).Methods("POST")

	fmt.Println("Listining port 8080")
	http.ListenAndServe(":8080", r)
}
