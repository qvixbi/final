package api

import (
	"fmt"
	"net/http"
	"time"

	"go_final_project/pkg/nextdate"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	nowStr := query.Get("now")
	date := query.Get("date")
	repeat := query.Get("repeat")

	now, err := time.Parse(nextdate.DateFormat, nowStr)
	if err != nil {
		http.Error(w, "неверная текущая дата", http.StatusBadRequest)
		return
	}

	result, err := nextdate.NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, result)
}
