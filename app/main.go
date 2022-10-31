package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var listen_port uint64

	flag.Uint64Var(&listen_port, "p", 5000, "Specify port to listen to")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./app/static")))

	var listener = ":" + strconv.FormatUint(listen_port, 10)
	log.Printf("Listening on port %s", listener)
	log.Fatal(http.ListenAndServe(listener, nil))
}
