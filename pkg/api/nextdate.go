package api

import (
	"fmt"
	"net/http"
	"time"

	"go_final_project/pkg/nextdate"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Получаем параметры запроса.
	nowStr := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	// Проверяем текущую дату.
	now, err := time.Parse("20060102", nowStr)
	if err != nil {
		http.Error(w, "Неверная текущая дата", http.StatusBadRequest)
		return
	}

	// Вычисляем следующую дату.
	result, err := nextdate.NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем дату обычным текстом.
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	fmt.Fprint(w, result)
}
