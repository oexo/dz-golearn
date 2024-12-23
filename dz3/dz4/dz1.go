package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

/*
Напишите программу, которая принимает строки через канал и подсчитывает количество слов в каждой строке. Используйте несколько горутин для обработки строк и один канал для передачи результатов.
Условно, на вход строка
"Всем привет!
Следующая лекция в среду
Увидимся на лекции!
Результат
Word count: 2
Word count: 4
Word count: 3

*/

// подсчитывает количество слов в строке

func wordCount(s string) int {
	count := 0
	inWord := false
	for _, rune := range s {
		if unicode.IsSpace(rune) || unicode.IsPunct(rune) {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}
	return count
}

func main() {
	c := make(chan string)
	defer close(c)

	text := ""
	strings := make([]string, 0)

	//бесконечно читаем канал
	go func() {
		for {
			strings = append(strings, <-c)
		}
	}()

	// бесконечно пишем в канал из stdin пок не получим слово "Result\n"
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ = reader.ReadString('\n')
		if text == "Result\n" {
			break
		}
		go func() { c <- text }()
	}

	fmt.Println(strings)

	// выводим результат
	fmt.Println("Результат:")
	for _, v := range strings {
		go func(v string) {
			fmt.Println("string:", v, "Word count:", wordCount(v))
		}(v)
	}
	time.Sleep(2 * time.Second)
}
