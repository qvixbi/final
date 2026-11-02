package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// db хранит подключение к базе данных.
var db *sql.DB

// schema содержит SQL-команды для создания таблицы и индекса.
const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '',
    title VARCHAR(255) NOT NULL DEFAULT '',
    comment TEXT NOT NULL DEFAULT '',
    repeat VARCHAR(128) NOT NULL DEFAULT ''
);

CREATE INDEX scheduler_date ON scheduler(date);
`

// Init открывает базу данных и при необходимости создаёт таблицу.
func Init(dbFile string) error {
	// Проверяем, существует ли файл базы данных.
	_, err := os.Stat(dbFile)

	install := false

	if os.IsNotExist(err) {
		install = true
	} else if err != nil {
		return fmt.Errorf("ошибка проверки файла БД: %w", err)
	}

	// Открываем базу данных.
	database, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("ошибка открытия БД: %w", err)
	}

	// Проверяем соединение.
	if err := database.Ping(); err != nil {
		database.Close()
		return fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	// Если база создаётся впервые, создаём таблицу и индекс.
	if install {
		if _, err := database.Exec(schema); err != nil {
			database.Close()
			return fmt.Errorf("ошибка создания схемы БД: %w", err)
		}
	}

	db = database

	return nil
}

// Close закрывает подключение к базе данных.
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
