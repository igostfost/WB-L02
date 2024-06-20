package repository

import (
	"calendar_api"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Repository struct {
	mu          sync.RWMutex
	eventsStore map[int]calendar_api.Event
}

func NewRepository() *Repository {
	return &Repository{
		eventsStore: make(map[int]calendar_api.Event),
	}
}

func (r *Repository) GetStoreLen() int {
	return len(r.eventsStore)
}

func (r *Repository) AddEventToStore(event calendar_api.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.eventsStore[event.ID]; exists {
		return fmt.Errorf("event with id %d already exists", event.ID)
	}

	r.eventsStore[event.ID] = event
	fmt.Println("Added event to store")
	return nil
}

func (r *Repository) DeleteEventFromStore(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exist := r.eventsStore[id]; !exist {
		return fmt.Errorf("event with id %d does not exist", id)
	}
	delete(r.eventsStore, id)
	fmt.Println("Deleted event from store")
	return nil
}

func (r *Repository) UpdateEventInStore(id int, name string, userId int, date time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	event, ok := r.eventsStore[id]
	if !ok {
		return fmt.Errorf("event with id %d not found", id)
	}
	if name != "" {
		// fmt.Println("поменяли name")
		event.Name = name

	}
	if !date.IsZero() {
		event.Date = date
	}
	if userId != 0 {
		event.UserID = userId
	}
	//обновили
	r.eventsStore[id] = event
	r.PrintMapStore()
	return nil
}

func (r *Repository) GetEventsForDayFromStore(date time.Time) ([]calendar_api.Event, error) {
	var result []calendar_api.Event
	date = date.Truncate(24 * time.Hour)

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, event := range r.eventsStore {
		eventDate := event.Date.Truncate(24 * time.Hour)
		if eventDate.Equal(date) {
			result = append(result, event)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no events found for day: %s", date.Format("2006-01-02"))
	}

	return result, nil
}

func (r *Repository) GetEventsForWeekFromStore(date time.Time) ([]calendar_api.Event, error) {

	var result []calendar_api.Event

	// Текущая дату обрезаем до начала дня
	startDate := date.Truncate(24 * time.Hour)
	endDate := startDate.AddDate(0, 0, 7) // Конечная дата через 7 дней

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, event := range r.eventsStore {
		eventDate := event.Date.Truncate(24 * time.Hour)
		if (eventDate.Equal(startDate) || eventDate.After(startDate)) && (eventDate.Equal(endDate) || eventDate.Before(endDate)) {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no events found for week: %s", date.Format("2006-01-02"))
	}

	return result, nil
}

func (r *Repository) GetEventsForMonthFromStore(firstDayOfMonth time.Time, lastDayOfMonth time.Time) ([]calendar_api.Event, error) {

	// Проверка на валидность входных данных
	if firstDayOfMonth.After(lastDayOfMonth) {
		return nil, errors.New("invalid date range: firstDayOfMonth is after lastDayOfMonth")
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []calendar_api.Event
	for _, event := range r.eventsStore {
		// Проверяем, находится ли дата события внутри месяца
		if (event.Date.Equal(firstDayOfMonth) || event.Date.After(firstDayOfMonth)) && event.Date.Before(lastDayOfMonth.AddDate(0, 0, 1)) {
			// Добавляем событие в результат
			result = append(result, event)
		}
	}

	return result, nil
}

func (r *Repository) PrintMapStore() {
	for key, value := range r.eventsStore {
		fmt.Printf("key is %d\n", key)
		fmt.Printf("id is %v\n", value.ID)
		fmt.Printf("name is %v\n", value.Name)
		fmt.Printf("userID is %v\n", value.UserID)
		fmt.Printf("date is %v\n", value.Date)
	}
}
