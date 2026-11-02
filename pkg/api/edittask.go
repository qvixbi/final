package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"go_final_project/pkg/db"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Читаем JSON из запроса.
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Проверяем идентификатор.
	if strings.TrimSpace(task.ID) == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	// Проверяем заголовок.
	if strings.TrimSpace(task.Title) == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан заголовок задачи",
		})
		return
	}

	// Проверяем дату и правило повторения.
	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Обновляем задачу в базе данных.
	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Возвращаем пустой JSON при успешном обновлении.
	writeJSON(w, struct{}{})
}
