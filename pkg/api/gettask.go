package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем идентификатор из URL.
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	// Получаем задачу из базы данных.
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Возвращаем найденную задачу.
	writeJSON(w, task)
}
