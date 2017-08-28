package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os/user"
)

type webData struct {
    name string
    a map[string]string
}

func main() {
    user, _ := user.Current()
    name := flag.String("username", user.Username, "Username used to lock")
    flag.Parse()
    fmt.Println("Using username: " + *name)
    x := map[string]string{"name": "john", "last_name": "smith"}
    w := webData{*name, x}
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
