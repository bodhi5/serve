package main

import (
  "flag"
  "net/http"
  "fmt"
)

func main() {

  dir := flag.String("d", ".", "Directory to serve, defaults to CWD")
  port := flag.String("p", "3000", "Port to serve on defaults to 3000")

  flag.Parse()

  fmt.Printf("Serving directory %v at http://0.0.0.0:%v.", *dir, *port)
  panic(http.ListenAndServe("0.0.0.0:" + *port, http.FileServer(http.Dir(*dir))))
}