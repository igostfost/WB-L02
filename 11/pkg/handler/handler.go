package handler

import (
	"calendar_api/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) InitRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.CreateEventHandler)
	mux.HandleFunc("/update_event", h.updateEventHandler)
	mux.HandleFunc("/delete_event", h.deleteEventHandler)
	mux.HandleFunc("/events_for_day", h.eventsForDayHandler)
	mux.HandleFunc("/events_for_week", h.eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", h.eventsForMonthHandler)

	// Добавляем middleware для логирования запросов
	handlerWithLogging := h.loggingMiddleware(mux)

	return handlerWithLogging
}
