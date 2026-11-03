package api

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/nextdate"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "метод не поддерживается",
		})
		return
	}

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "не указан идентификатор",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, db.ErrTaskNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			if errors.Is(err, db.ErrTaskNotFound) {
				writeJSON(w, http.StatusNotFound, map[string]string{
					"error": err.Error(),
				})
				return
			}

			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, struct{}{})
		return
	}

	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := db.UpdateDate(next, id); err != nil {
		if errors.Is(err, db.ErrTaskNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, struct{}{})
}
