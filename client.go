package main

import (
	"bases31/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL_GET    string = "http://localhost:8080/get"
	URL_POST   string = "http://localhost:8080/post"
	URL_PUT    string = "http://localhost:8080/put"
	URL_DELETE string = "http://localhost:8080/delete"
)

func main() {

	new := storage.Post{
		ID:         3,
		Title:      "Новая заметка",
		Content:    "Про вашего мальчика",
		AuthorID:   3,
		AuthorName: "Печкин",
		CreatedAt:  "11:00:00",
	}
	requestBody, err := json.Marshal(new)

	//GET
	/*
		resp, err := http.Get(URL_GET)
		if err != nil {
			fmt.Println("Ошибка при отправке GET-запроса:", err)
			return
		}
	*/
	// POST - добавление new

	resp, err := http.Post(URL_POST, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}

	// UPDATE - изменение записи по ID экземпляра new
	/*
		resp, err := http.Post(URL_PUT, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println(err)
		}
	*/
	// DELETE - удаление записи по ID экземпляра new
	/*
		resp, err := http.Post(URL_DELETE, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println(err)
		}
	*/

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Сервер вернул ошибку:", resp.Status)
		return
	}

	defer resp.Body.Close() // Закрываем тело ответа при завершении функции
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ сервера:", string(responseBytes))
}
