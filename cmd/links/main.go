package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
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
    fmt.Println(postToWeb("/post", w))
}

func postToWeb(endpoint string, data webData) string {

    fmt.Println("In function: ", data.name)
    _url := "http://localhost:5000" + endpoint

    params := url.Values{}
    for k, v := range data.a {
        params.Add(k, v)
    }

    resp, err := http.PostForm(_url, params)

    if err != nil {
        //fmt.Println("Error: ", err)
        panic(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return string(body)
}
