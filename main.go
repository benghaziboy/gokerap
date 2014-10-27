package main

import (
    "encoding/json"
    "fmt"
    "log"
    "log/syslog"
    "net/http"
    "gokerap/rapper"
)

type Route struct {
    Name         string `json:"name"`
    Description string `json:"description"`
}

func handleIndex(w http.ResponseWriter, request *http.Request) {
    r := &Route{"/rappers", "GET a list of rappers."}
    content, _ := json.Marshal(r)
    w.Write(content)
}

func handleRappers(rwriter http.ResponseWriter, request *http.Request) {
    guc := &rapper.Rapper{"Gucci Mane"}
    json.NewEncoder(rwriter).Encode(guc)
    rwriter.Header().Set("Content-Type:", "application/json")
}   

func main() {
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "go_log")
    if e != nil {
        fmt.Println("Hey Syslog")
        log.SetOutput(logwriter)
    }
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/rappers", handleRappers)
    http.ListenAndServe(":8080", nil)
}
