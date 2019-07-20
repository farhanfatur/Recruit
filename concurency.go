package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var get = func() string {
		var url, err = http.Get("http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/profilGet?museum_id=4A33CF6F-A284-4E42-830B-E7DC755614CD")
		if err != nil {
			log.Fatal(err.Error())
		}
		resp, err := ioutil.ReadAll(url.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		return string(resp)
	}

	go get()
	fmt.Println(get())
}
