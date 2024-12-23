package main

import (
	"fmt"
	"time"
)

/*
Описание задачи:
Реализуйте простую модель многопользовательского чата, где несколько пользователей могут отправлять сообщения в общий канал. Каждое сообщение должно содержать имя отправителя и текст сообщения. Создайте несколько горутин для имитации пользователей, которые отправляют сообщения.

Примерный вывод:
[User3]: Message 1 from User3
[User1]: Message 1 from User1
[User2]: Message 1 from User2
[User2]: Message 2 from User2


*/

func main() {
	fmt.Println("vim-go")
	c := make(chan string)
	defer close(c)

	text := ""

	//бесконечно читаем канал
	go func() {
		for {
			fmt.Println("Входящее сообщение:", <-c)
		}
	}()

	// бот номер два
	go func() {
		for i := 0; i < 10; i++ {
			text = fmt.Sprintf("%s%d", "[User2]: Message # ", i)
			c <- text
			time.Sleep(1 * time.Second)
		}
	}()
	// бот номер один
	for i := 0; i < 10; i++ {
		text = fmt.Sprintf("%s%d", "[User1]: Message # ", i)
		go func() { c <- text }()
		time.Sleep(1 * time.Second)
	}

}
