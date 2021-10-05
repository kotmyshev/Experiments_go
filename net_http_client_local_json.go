// test json client
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"os"
	"bytes"
	//"net/url"
	"time"
)

const (
	n = 42
)

func main() {

	fmt.Println("Hello backend world!", n)

	client := http.Client{}
	client.Timeout = time.Second * 4

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

	bo := BookData{"Фрэнк", "Герберт", "Дюна", "978-5-17-087261-9", 1965, "Фантастика", "Новая", "Нейштадтская литературная премия", "Твердая", "Английский", true, true}

	bw := BookCategory{[]string{"История", "Триллер", "Ужасы"}, []string{"Новая", "Мятая"}, []string{"Букер", "Ясная Поляна"}, []string{"Твердая"}, []string{"Японский", "Суахили"}}

	useraddr := PostAddress{"Иванов", "Иван", "Феоктистович", "443022", "Самара", "Проспект Фридриха Энгельса", "28Г", "К1", "34", true}

	//fmt.Println(bd)

	m := ExchangeMessage{bo, bw, useraddr}

	b, err := json.Marshal(m)

	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Post("http://localhost:8080", "application/json", bytes.NewBuffer(b))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode)
	//fmt.Println(string(b))

	// file, err := os.Create("textfile.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// file.Write(b)

}
