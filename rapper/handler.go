package rapper

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type NewRapperForm struct {
    Name    string  `json:"name"`  
}

func RapperHandler(w http.ResponseWriter, request *http.Request) error {

        if request.Method == "POST" {
            var args NewRapperForm
            err := json.NewDecoder(request.Body).Decode(&args)
            fmt.Println(args)

            if err != nil {
                    return err
            }
            defer request.Body.Close()

            r, err := New(args.Name)
            if err != nil {
                return err
            }

            w.WriteHeader(http.StatusCreated)
            return json.NewEncoder(w).Encode(r)
        } 

        if request.Method == "GET" {
            r, err := List()
            
            if err != nil {
                return err
            }

            w.WriteHeader(http.StatusOK)
            return json.NewEncoder(w).Encode(r)
        }

        http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
        return nil
}
