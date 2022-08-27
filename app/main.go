package main

import (
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./app/static")))

    log.Fatal(http.ListenAndServe(":8081", nil))
}
