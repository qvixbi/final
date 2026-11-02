package api

import (
	"net/http"
	"strings"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/nextdate"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, map[string]string{
			"error": "Метод не поддерживается",
		})
		return
	}

	// Получаем идентификатор задачи.
	id := r.URL.Query().Get("id")

	if strings.TrimSpace(id) == "" {
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

	// Если задача одноразовая, удаляем её.
	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, struct{}{})
		return
	}

	// Для периодической задачи рассчитываем следующую дату.
	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Сохраняем новую дату в базе данных.
	if err := db.UpdateDate(next, id); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, struct{}{})
}
