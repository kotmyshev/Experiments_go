//server1
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func printHello(wr http.ResponseWriter) {
	message := []byte("Hello, web!")
	_, err := wr.Write(message)
	if err != nil {
		log.Fatal(err)
	}

}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	printHello(w)
}

func reqHandler(w http.ResponseWriter, req *http.Request) {
	m := req.Method

	switch m {
	case "POST":
		fmt.Println("POST")
		fmt.Println(req.Header)
		fmt.Println(req.Body)
		req.ParseForm()

		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		v := req.Form
		h := v.Get("hello")
		l := v.Get("life")
		fmt.Println(h)
		fmt.Println(l)

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
		mess := []byte(bytesRepresentation)
		_, errq := w.Write(mess)
		if errq != nil {
			log.Fatal(errq)
		}

	case "GET":
		fmt.Println("GET")
		fmt.Println(req.Header)
		fmt.Println(req.Body)
	}
}

func main() {

	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/", reqHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
