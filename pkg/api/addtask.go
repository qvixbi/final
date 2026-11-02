package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/nextdate"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Читаем JSON из запроса
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Проверяем заголовок
	if strings.TrimSpace(task.Title) == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверяем дату и правило повторения
	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Добавляем задачу в базу
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Возвращаем ID созданной задачи
	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}

func checkDate(task *db.Task) error {
	now := time.Now()

	// Если дата не указана, используем сегодняшний день
	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	// Проверяем формат и существование даты
	_, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}

	// Проверяем правило повторения, если оно указано
	var next string
	if task.Repeat != "" {
		next, err = nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	// Если дата задачи раньше сегодняшней
	if task.Date < now.Format("20060102") {
		if task.Repeat == "" {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}

	return nil
}
