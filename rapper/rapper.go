package rapper

import (
    "gokerap/db"
    "log"
)

type Rapper struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func Get(id int) (*Rapper, error) {
    r := Rapper{}
    err := db.Conn.QueryRow(getRapper, id).Scan(
        &r.Id, &r.Name)

    return &r, err
}

func List() ([]Rapper, error) {
    rappers := make([]Rapper, 0)

    rows, err:= db.Conn.Query(listRapper)
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        r := Rapper{}
        if err = rows.Scan(&r.Id, &r.Name); err != nil {
            log.Fatal(err)
        } else {
            rappers = append(rappers, r)
        }
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
        return nil, err
    }

    return rappers, err
}

func New(name string) (*Rapper, error) {
    r := &Rapper {
        Name: name,
    }

    err := db.Conn.QueryRow(createRapper, r.Name).Scan(&r.Id)
    if err != nil {
        return nil, err
    }

    return r, err
}

func init() {
    _, err := db.Conn.Exec(createRapperTable)
    if err != nil {
        log.Println(err)
    }
}
