package main

import (
  "ending/json"
  "log"
  "net/http"
  "time"
)

func apiClockHandler(w http.ResponseWriter, r *http.Request) {
  
  type ResponseBody struct {
    Time time.Time `json:"time"`
  }
  
  rb := &ResponseBody {
    Time: time.Now();
  }
  
  w.Header().Set("Content-type", "application/json")
  
  if err := json.NewEncoder(w).Encode(rb); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/api/clock", apiClockHandler)
  
  log.Fatal(http.ListenAndServe(":8080", nil))
}

