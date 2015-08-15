package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port = 80
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("invalid port %s\n", os.Args[1])
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	portStr := ":" + strconv.Itoa(port)
	err = http.ListenAndServe(portStr, http.FileServer(http.Dir(wd)))
	if err != nil {
		log.Fatal(err)
	}
}
