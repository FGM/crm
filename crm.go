package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/FGM/crm/api"
	"github.com/spf13/viper"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MustKey() string {
	viper.SetDefault("api_key", api.DEFAULT_KEY)
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.crm")
	err := viper.ReadInConfig()
	check(err)

	key := viper.Get("api_key").(string)
	if key == api.DEFAULT_KEY {
		log.Fatalln("Could not read a usable API key in config.")
	}

	return key
}

func main() {
	var err error
	var key string = MustKey()

	log.Printf("Using API key %s\n", key)

	var user *url.Userinfo

	user = url.User(key)
	check(err)

	var url *url.URL
	url, err = url.Parse(api.BASE_URL + "Contacts")
	check(err)
	url.User = user

	var response *http.Response

	response, err = http.Get(url.String())
	defer response.Body.Close()
	check(err)
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Request error %s\n", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println(string(body))
}
