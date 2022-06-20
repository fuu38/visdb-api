package main

import (
	"log"
	"net/http"
	resolver "visdb/resolver"
)

func main() {
	http.HandleFunc("/", resolver.ShowAll)
	http.HandleFunc("/insert", resolver.Insert)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
