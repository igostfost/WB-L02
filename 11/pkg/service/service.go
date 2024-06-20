package service

import (
	"calendar_api"
	"calendar_api/pkg/repository"
	"fmt"
	"time"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) AddEvent(newEvent calendar_api.Event) error {
	//eventId := s.Repo.GetStoreLen()
	//newEvent := calendar_api.Event{ID: eventId, Name: name, Date: date, UserID: userID}
	// return nil
	return s.Repo.AddEventToStore(newEvent)
}

//func (s *Service) GetEventByID(id int) (event calendar_api.Event, ok bool) {
//	event, ok = s.Repo.GetEventByIDFromStore(id)
//	return event, ok
//}

func (s *Service) UpdateEvent(event calendar_api.Event) error {
	fmt.Println("name in services")
	return s.Repo.UpdateEventInStore(event.ID, event.Name, event.UserID, event.Date)
}
func (s *Service) DeleteEvent(id int) error {
	return s.Repo.DeleteEventFromStore(id)
}

func (s *Service) GetEventsForDay(date time.Time) ([]calendar_api.Event, error) {
	return s.Repo.GetEventsForDayFromStore(date)
}

func (s *Service) GetEventsForWeek(date time.Time) ([]calendar_api.Event, error) {
	return s.Repo.GetEventsForWeekFromStore(date)
}

func (s *Service) GetEventsForMonth(firstDayOfMonth time.Time, lastDayOfMonth time.Time) ([]calendar_api.Event, error) {
	return s.Repo.GetEventsForMonthFromStore(firstDayOfMonth, lastDayOfMonth)
}
