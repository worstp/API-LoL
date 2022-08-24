package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetItems() {
	res, err := http.Get("https://perodriguezl-league-of-legends-v1.p.rapidapi.com/lol/items")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}