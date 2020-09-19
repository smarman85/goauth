package main

import (
  "encoding/json"
  "net/http"
  "log"
)

type person struct {
  First string
}

func main() {
  // create server on 8087
  http.HandleFunc("/encode", foo)
  http.HandleFunc("/decode", bar)
  http.ListenAndServe(":8087", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  // curl localhost:8087/encode
  p1 := person{
    First: "Sean",
  }

  err := json.NewEncoder(w).Encode(&p1)
  if err != nil {
    log.Println("Encoded bad data", err)
  }
}

func bar(w http.ResponseWriter, r *http.Request) {
  //curl -XGET -H "Content-type: application/json" -d '{"First": "Tiffany"}' 'localhost:8087/decode'
  var p1 person
  err := json.NewDecoder(r.Body).Decode(&p1)
  if err != nil {
    log.Println("Decoded bad data", err)
  }
  log.Println("Person:", p1)
}
