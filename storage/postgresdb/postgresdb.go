package postgresdb

import (
	"bases31/storage"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func InitDb() *Storage {

	// Строка подключения к базе данных PostgreSQL.
	// Замените данными вашей базы данных.
	connStr := "user=leka password=serpent dbname=mydb host=localhost sslmode=disable"

	// Открываем соединение с базой данных.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Проверяем подключение к базе данных.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Успешное подключение к PostgreSQL!")

	// Далее вы можете выполнять запросы к базе данных.

	storage := Storage{
		db: db,
	}
	return &storage
}

func (s *Storage) Posts() ([]storage.Post, error) {

	rows, err := s.db.Query("SELECT id,author_id,title,content,created FROM posts")
	if err != nil {
		log.Fatal(err)

	}
	defer rows.Close()

	// Обрабатываем результаты запроса
	var post []storage.Post
	for rows.Next() {
		var row storage.Post

		err := rows.Scan(&row.ID, &row.AuthorID, &row.Title, &row.Content, &row.CreatedAt)
		if err != nil {
			log.Fatal(err)
			fmt.Println("________")
		}
		post = append(post, row)
	}
	return post, nil
}

func (s *Storage) AddPost(stor storage.Post) error {
	// Выполняем SQL-запрос для вставки данных
	_, err := s.db.Exec("INSERT INTO posts (author_id, title, content, created) VALUES ($1, $2, $3, $4)", stor.AuthorID, stor.Title, stor.Content, stor.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Println("BD Postgress : Данные добавлены")
	return nil
}
func (s *Storage) UpdatePost(stor storage.Post) error {
	// Выполняем SQL-запрос для обновления данных
	_, err := s.db.Exec("UPDATE posts SET title = $1 WHERE id = $2", "Вносим пропаганды", stor.AuthorID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BD Postgress : Данные обновлены")
	return nil
}
func (s *Storage) DeletePost(stor storage.Post) error {
	_, err := s.db.Exec("DELETE FROM posts where id=$1", stor.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BD Postgress : Данные удалены")
	return nil
}
