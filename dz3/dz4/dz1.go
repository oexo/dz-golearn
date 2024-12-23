package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

var String chan string

func PrintFromChannel(c chan string, t *string) {
	fmt.Println(*t)
	c <- *t
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(String)

	c := make(chan string)
	fmt.Println(c)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	go PrintFromChannel(c, &text)
	time.Sleep(2 * time.Second)
	fmt.Println(c)

	text2 := <-c
	fmt.Println(text2)

}
