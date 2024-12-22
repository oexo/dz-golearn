package main

import "fmt"

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

func main() {
	fmt.Println("vim-go")
	fmt.Println(String)

	c := make(chan string)
	fmt.Println(c)
}
