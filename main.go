package main

import (
    "flag"
    "github.com/gorilla/mux"
    "log"
    "log/syslog"
    "net/http"
    "gokerap/rapper"
    "gokerap/user"
)

var (
    port = flag.String("port", ":8080", "HTTP port to listen on")
)

func Router() *mux.Router{
    router := mux.NewRouter()
    router.Handle("/api/v1/user/new", AnonymousApiHandler(user.UserRegistrationHandler)).Name("user_registration")
    router.Handle("/api/v1/user/login", AnonymousApiHandler(user.UserAuthHandler)).Name("user_authenticate")
    router.Handle("/api/v1/user/{user_id:[0-9]+}", AuthApiHandler(user.UserIdHandler)).Name("user_id")
    router.Handle("/api/v1/rapper/", AnonymousApiHandler(rapper.RapperHandler)).Name("rapper")

    return router
}

func main() {
    logwriter, _ := syslog.New(syslog.LOG_NOTICE, "go_log")
    http.Handle("/", Router())

    if err := http.ListenAndServe(*port, nil); err != nil {
        log.SetOutput(logwriter)
        log.Fatal(err)
    }
}
