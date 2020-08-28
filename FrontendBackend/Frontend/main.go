package main

import (
	"time"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	// loops every 15 seconds
	for {
		// make a sample HTTP GET request
		url := "http://localhost:8080"
		send_url := url + "/send" 

		requestBody := strings.NewReader(`
			{
				"name":"test",
				"salary":"123",
				"age":"23"
			}
		`)
		
		// post some data
		res, err := http.Post(
			send_url,
			"application/json; charset=UTF-8",
			requestBody,
		)

		// check for response error
		if err != nil {
			log.Fatal( err )
		}

		// read all response body
		data, _ := ioutil.ReadAll( res.Body )
		dataStr := string(data)

		// close response body
		res.Body.Close()

		// print `data` as a string
		logger.Println( "%s\n", dataStr )
		time.Sleep(15 * time.Second)
	}
}
