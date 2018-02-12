package main

import (
    "encoding/json"
    "fmt"
    "strconv"
)

type Product struct {
    Name       string
    Price      string
    PriceFloat float64
}

func main() {
    s := `{"name":"Galaxy Nexus", "price":"3460.00"}`
    var pro Product
    err := json.Unmarshal([]byte(s), &pro)
    if err == nil {
        pro.PriceFloat, err = strconv.ParseFloat(pro.Price, 64)
        if err != nil { fmt.Println(err) }
        fmt.Printf("%+v\n", pro)
    } else {
        fmt.Println(err)
        fmt.Printf("%+v\n", pro)
    }
}
