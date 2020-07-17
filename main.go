package main

import (
  "log"
  "net/http"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FileServer(http.Dir("/etc/gohttpserver/doc")))))
  log.Fatal(http.ListenAndServe(":8080", nil))
}
/*

func main() {
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(
        http.Dir("/etc/gohttpserver/doc"),
      ),
    ),
  )

  log.Fatal(
    http.ListenAndServe(":8080", nil),
  )
}

*/

