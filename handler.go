package main

import (
	"log"
	"net/http"
	"gokerap/app"
	"gokerap/user"
)

type AnonymousApiHandler func(w http.ResponseWriter, r *http.Request) error

func (a AnonymousApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := a(w, r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type AuthApiHandler func(w http.ResponseWriter, r *http.Request, c *app.Context) error

func (a AuthApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := user.AuthenticateToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := app.Context{
		UserId: u.Id,
	}

	err = a(w, r, &c)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
