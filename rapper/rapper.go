package rapper

type Rapper struct {
    Name string `json:"name"`
}

func New(name string) Rapper {
    rapper := Rapper{name}
    return rapper
}
