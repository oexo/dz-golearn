package main

/*
ДЗ №5
Цель: Создать HTTP-сервер на языке Go, который будет обрабатывать заявки студентов на поступление в университет. Сервер должен принимать данные о студентах, проверять их баллы и выводить список поступивших студентов.

Задачи:
Создание структуры данных:
Определите структуру Student, которая будет содержать следующие поля:
FullName (строка) — полное имя студента.
MathScore (целое число) — балл по математике.
InformaticsScore (целое число) — балл по информатике.
EnglishScore (целое число) — балл по английскому языку.
Создание HTTP-сервера:
Реализуйте HTTP-сервер, который будет слушать на порту 8080.
Обработчик для поступления:
Создайте обработчик для POST-запросов на маршрут /apply, который будет принимать JSON с данными студента.
В обработчике проверьте, если сумма баллов по трем предметам (математика, информатика, английский) больше или равна 14, то добавьте студента в список поступивших. В противном случае, верните сообщение о том, что студент не поступил.
Создание студентов:
Создайте трех студентов (клиентов) с заранее определенными баллами:
Два студента должны иметь общую сумму баллов >= 14.
Один студент должен иметь общую сумму баллов < 14.
Обработчик для вывода поступивших студентов:
Создайте новый маршрут /admitted, который будет возвращать список всех студентов, которые поступили. Список должен быть представлен в формате JSON.

*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Student struct {
	FullName         string `json:"name"`
	MathScore        int    `json:"math"`
	InformaticsScore int    `json:"informatics"`
	EnglishScore     int    `json:"english"`
}

type Students struct {
	Students []Student
}

var students Students

func postApply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	student := &Student{}

	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("got student:", student)
	w.WriteHeader(http.StatusCreated)
	if student.MathScore+student.InformaticsScore+student.EnglishScore < 14 {
		fmt.Println(student.FullName, "not accepted")
	} else {
		fmt.Println(student.FullName, "accepted")
		students.Students = append(students.Students, *student)
	}
}

func getAdmitted(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	for _, person := range students.Students {
		jsonData, err := json.Marshal(person)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))

	}
}

func main() {
	http.HandleFunc("/apply", postApply)
	http.HandleFunc("/admitted", getAdmitted)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
