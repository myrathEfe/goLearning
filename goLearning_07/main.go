package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc(path, handlerFunc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Merhaba bu benim ilk API sunucum")
	})
	fmt.Println("Sunucu başlatıldı : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
