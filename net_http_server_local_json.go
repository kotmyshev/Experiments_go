package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



type BookData struct {
	AutorFirstName string // Фамилия автора
	AutorLastName  string // Имя автора
	BookName       string // Название книги
	BookISBN       string // ISBN книги
	BookPubYear    int    // Год публикации книги
	Genre          string // Жанр
	Condition      string // Состояние книги (новая, в хорошем состоянии, б/у, потрепана)
	Laureate       string // Лауреат премии (название премии)
	Cover          string // Обложка (суперобложка, жесткая, мягкая, без обложки)
	Language       string //Язык
	Filming        bool   //была ли экранизирована книга (да, нет);
	RareBook       bool   // издание подарочное или ограниченного тиража;
}

type BookCategory struct {
	Genre     []string // Жанры желаемые
	Condition []string // Состояния книги
	Laureate  []string // Лауреаты каких премий
	Cover     []string // Типы обложек
	Language  []string // Языки на каоторых издана
}

type PostAddress struct {
	LastName      string //Фамилия получателя
	FirstName     string //Имя получателя
	Patronymic    string // Отчество получателя
	AddrIndex     string // Индекс почтовый
	AddrCity      string // Город
	AddrSreet     string // Улица
	AddrHouse     string // Дом
	AddrStructure string // Строение, корпус
	AddrApart     string // Квартира

	IsDefault bool // Адрес по умолчанию
}

type ExchangeMessage struct {
	Offer   BookData
	Wish    BookCategory
	Address PostAddress
}

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
		decoder := json.NewDecoder(req.Body)

		var m ExchangeMessage

		err := decoder.Decode(&m)

		if err != nil {
			log.Fatalln(err)
		}

		//fmt.Println(req.Body)

		du := m.Wish.Genre[0]

		fmt.Println(du)


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
