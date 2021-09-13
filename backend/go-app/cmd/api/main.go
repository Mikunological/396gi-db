package main

import (
    "fmt"
    "net/http"
    "github.com/github.com/Mikunological/396gi-db/app/interfaces/api"
)

func main() {
    port := 8000
    
    s := api.NewServer()
    if err := s.Init(); err != nil {
        fmt.Println(err)
    }
    s.Run(port)

    // // controller
    // http.HandleFunc("/", echoHello)
    // // port
    // http.ListenAndServe(":8000", nil)
}

