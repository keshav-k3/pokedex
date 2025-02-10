package models

type Pokemon struct {
    ID        int      `json:"id"`
    Name      string   `json:"name"`
    Types     []string `json:"types"`
    Height    int      `json:"height"`
    Weight    int      `json:"weight"`
    ImageURL  string   `json:"image_url"`
    Stats     Stats    `json:"stats"`
}

type Stats struct {
    HP        int `json:"hp"`
    Attack    int `json:"attack"`
    Defense   int `json:"defense"`
    SpAttack  int `json:"sp_attack"`
    SpDefense int `json:"sp_defense"`
    Speed     int `json:"speed"`
}