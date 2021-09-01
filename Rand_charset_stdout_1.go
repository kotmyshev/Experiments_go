// Rnd project 
// Rand_charset_stdout_1.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Генератор псевдослучайных буквенных последовательностей", "\n", "Введите желаемое количеcтво букв:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода bufio.NewReader")
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	ncount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка ввода:", input, "- не целое число!")
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix() + int64(ncount))

	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder

	for i := 0; i < ncount; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	fmt.Println(output.String())

	fmt.Println("Ок!")
}
