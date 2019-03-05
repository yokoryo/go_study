//URLを引数に指定し、その中身を表示
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	//url := flag.String("url", "http://localhost:8000", "URL")
	url := flag.String("url", "https://golang.org", "URL")
	flag.Parse()

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(body))
}