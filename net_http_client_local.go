// client project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	client := http.Client{}
	client.Timeout = time.Second * 2
	//response, err := client.Get("http://localhost:8080")
	//request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	response, err := client.Post("http://localhost:8080", "lala", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode)

}
