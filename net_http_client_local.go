// client project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bytes"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	client := http.Client{}
	client.Timeout = time.Second * 4
	//response, err := client.Get("http://localhost:8080")
	//request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	//r := strings.NewReader("Hello, Reader!!!!!!!!!!!!!!!dv")

	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}


	//response, err := client.Post("http://localhost:8080", "text/plain", bytes.NewBufferString("Hello, Reader!!!!!!!!!!!!!!!dv"))
	response, err := client.Post("http://localhost:8080", "application/json", bytes.NewBuffer(bytesRepresentation))
	//client.PostForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode)

}
