// client project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"net/url"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	client := http.Client{}
	client.Timeout = time.Second * 4
	//response, err := client.Get("http://localhost:8080")
	//request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	//r := strings.NewReader("Hello, Reader!!!!!!!!!!!!!!!dv")

	v := url.Values{}
	v.Set("hello", "Говно")
	v.Add("life", "Какашка")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	//response, err := client.Post("http://localhost:8080", "text/plain", bytes.NewBufferString("Hello, Reader!!!!!!!!!!!!!!!dv"))
	//response, err := client.Post("http://localhost:8080", "application/json", bytes.NewBuffer(bytesRepresentation))
	response, err := client.PostForm("http://localhost:8080", v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode)
	//fmt.Println(response.Write())
	//fmt.Println(response.Body)

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
	log.Println(result["hello"])
}
