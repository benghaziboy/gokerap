package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"gokerap/app"
)

type NewUserForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegistrationHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	var args NewUserForm
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	u, err := New(args.Email, args.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(u)
}

func UserAuthHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	var args NewUserForm
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	u, err := AuthenticatePassword(args.Email, args.Password)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(u)
}

func UserHandler(w http.ResponseWriter, r *http.Request, c *app.Context) error {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	u, err := Find(c.UserId)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(u)
}

func UserIdHandler(w http.ResponseWriter, r *http.Request, c *app.Context) error {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	u, err := Find(id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(u)
}
