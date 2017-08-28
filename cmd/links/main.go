package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    s := "sijis"
    fmt.Println(s)

    fmt.Println(postToWeb("/"))
}

func postToWeb(endpoint string) string {

    url := "http://localhost:5000" + endpoint
    resp, err := http.Get(url)
    if err != nil {
        //fmt.Println("Error: ", err)
        panic(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return string(body)
}
