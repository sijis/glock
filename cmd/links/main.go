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
    dataset  map[string]string
}

var (
    endpoint = "http://localhost:5000/locker"
)

func main() {
    user, _ := user.Current()
    name := flag.String("username", user.Username, "Username used to lock/unlock chest.")
    action := flag.String("action", "locked", "Action to take.")
    chest := flag.String("chest", "", "Which chest to lock/unlock.")
    flag.Parse()

    dataset := map[string]string{
        "username": "@" + *name,
        "action":   *action,
        "chest":    *chest,
    }

    w := webData{dataset}
    fmt.Println(postToWeb(w))
}

func postToWeb(data webData) string {

    fmt.Println("[DEBUG] postToWeb: ", data)
    params := url.Values{}
    for k, v := range data.dataset {
        params.Add(k, v)
    }

    resp, err := http.PostForm(endpoint, params)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}
