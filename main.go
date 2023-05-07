package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	dir, _ := os.Getwd()
	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir(dir))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
