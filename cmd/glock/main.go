package glock

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/user"

	"github.com/spf13/viper"
)

type webData struct {
	dataset map[string]string
}

// Run Entry point in launching cli
func Run() {
	user, _ := user.Current()
	name := flag.String("username", user.Username, "Username used to lock/unlock chest.")
	action := flag.String("action", "locked", "Action to take.")
	chest := flag.String("chest", "", "Which chest to lock/unlock.")
	flag.Parse()

	viper.SetConfigName("glock-config")
	viper.AddConfigPath("$HOME/.config/")

	if viper.ReadInConfig() != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

	viper.SetDefault("endpoint", "http://127.0.0.1:5000")
	endpoint := viper.GetString("endpoint")

	dataset := map[string]string{
		"username": "@" + *name,
		"action":   *action,
		"chest":    *chest,
	}

	w := webData{dataset}
	fmt.Println(postToWeb(w, endpoint))
}

func postToWeb(data webData, endpoint string) string {

	//fmt.Println("[DEBUG] postToWeb: ", data)
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
