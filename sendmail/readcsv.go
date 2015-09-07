package main

import (
    "github.com/cuixin/csv4g"
    "fmt"
)

type Test struct {
    Id   int
    Name string
    Email string
}

func main() {
    csv, err := csv4g.New("./email.csv", ',', Test{}, 1)
    if err != nil {
        fmt.Errorf("Error %v\n", err)
        return
    }
    for i := 0; i < csv.LineLen; i++ {
        tt := &Test{}
        err = csv.Parse(tt)
        if err != nil {
            fmt.Printf("Error on parse %v\n", err)
            return
        }
        fmt.Println(tt)
        fmt.Printf("email:%s\n",tt.Email)
    }
}
