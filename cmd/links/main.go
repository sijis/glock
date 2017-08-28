package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    s := "sijis"
    fmt.Println(s)

    resp, err := http.Get("http://localhost:5000")
    if err != nil {
        //fmt.Println("Error: ", err)
        panic(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
