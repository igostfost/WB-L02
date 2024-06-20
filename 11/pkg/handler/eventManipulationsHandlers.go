package handler

import (
	"calendar_api"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func respondWithJSON(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	var event calendar_api.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "Invalid request payload"})
		return
	}

	err := h.service.AddEvent(event)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to add event"})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{Result: fmt.Sprintf("Event added")})
}

func (h *Handler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	var event calendar_api.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "Invalid request payload"})
		return
	}

	err := h.service.UpdateEvent(event)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to add event"})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{Result: fmt.Sprintf("Event updated")})

}

func (h *Handler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	var event calendar_api.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "Invalid request payload"})
		return
	}

	err := h.service.DeleteEvent(event.ID)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to add event"})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{Result: fmt.Sprintf("Event deleted")})
}

func (h *Handler) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	// Получаем значения параметров из строки запроса
	queryValues := r.URL.Query()

	// Получаем значение параметра "date"
	dateString := queryValues.Get("date")

	// Проверяем, не является ли значение параметра "date" пустым
	if dateString == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "date parameter is missing"})
		return
	}

	// Преобразуем значение параметра "date" в формат time.Time
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "failed to parse date parameter"})
		return
	}

	// Получаем события для указанной даты через сервис
	events, err := h.service.GetEventsForDay(parsedDate)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to get events for day"})
		return
	}

	// Возвращаем успешный ответ с найденными событиями
	respondWithJSON(w, http.StatusOK, Response{Result: events})
}

func (h *Handler) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	// Получаем значения параметров из строки запроса
	queryValues := r.URL.Query()

	// Получаем значение параметра "date"
	dateString := queryValues.Get("date")

	// Проверяем, не является ли значение параметра "date" пустым
	if dateString == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "date parameter is missing"})
		return
	}

	// Преобразуем значение параметра "date" в формат time.Time
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "failed to parse date parameter"})
		return
	}

	//пока дата не станет понедельником текущей недели, вычитаем по дню
	for parsedDate.Weekday() != time.Monday {
		parsedDate = parsedDate.Add(-24 * time.Hour)
	}

	// Получаем события для указанной даты через сервис
	events, err := h.service.GetEventsForWeek(parsedDate)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to get events for week"})
		return
	}

	// Возвращаем успешный ответ с найденными событиями
	respondWithJSON(w, http.StatusOK, Response{Result: events})
}

func (h *Handler) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	// Получаем значения параметров из строки запроса
	queryValues := r.URL.Query()

	// Получаем значение параметра "date"
	dateString := queryValues.Get("date")

	// Проверяем, не является ли значение параметра "date" пустым
	if dateString == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "date parameter is missing"})
		return
	}

	// Преобразуем значение параметра "date" в формат time.Time
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{Error: "failed to parse date parameter"})
		return
	}

	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	fmt.Println(firstDayOfMonth)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	fmt.Println(lastDayOfMonth)
	res, err := h.service.GetEventsForMonth(firstDayOfMonth, lastDayOfMonth)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{Error: "Failed to get events for week"})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{Result: res})
}
