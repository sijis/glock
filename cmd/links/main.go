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
    endpoint string
    dataset map[string]string
}

func main() {
    user, _ := user.Current()
    name := flag.String("username", user.Username, "Username used to lock")
    flag.Parse()
    dataset := map[string]string{"name": *name, "misc": "other_stuff"}
    w := webData{"/post", dataset}
    fmt.Println(postToWeb(w))
}

func postToWeb(data webData) string {

    fmt.Println("[DEBUG] postToWeb: ", data)
    _url := "http://localhost:5000" + data.endpoint

    params := url.Values{}
    for k, v := range data.dataset {
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
