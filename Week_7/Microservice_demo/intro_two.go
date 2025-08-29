package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
        "github.com/gorilla/mux"
)

func main() {
        router := mux.NewRouter()
        router.HandleFunc("/resources", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "/resources")
        })
        router.HandleFunc("/resources/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "/resources/{id:[0-9]+}: %s!", mux.Vars(r)["id"])
        })
        router.HandleFunc("/resources/{id:[0-9]+}/values", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "/resources/{id:[0-9]+}/values: %s", mux.Vars(r)["id"])
        })
        srv := &http.Server{
                Handler:      router,
                Addr:         ":8080",
                WriteTimeout: 15 * time.Second,
                ReadTimeout:  15 * time.Second,
        }

        log.Fatal(srv.ListenAndServe())
}
