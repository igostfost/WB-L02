package main

import "fmt"

/* Реализуем мини систему домашнего кинотеатра. Наш фасад запускает просмотр фильма и заканчивает,
скрывая всю реализацию от клиента внутри себя
*/

type DVDPlayer struct {
}

func (dvd *DVDPlayer) On() {
	fmt.Println("DVDPlayer On")
}

func (dvd *DVDPlayer) Off() {
	fmt.Println("DVDPlayer Off")
}

type Projector struct{}

func (projector *Projector) On() {
	fmt.Println("Projector On")
}
func (projector *Projector) Off() {
	fmt.Println("Projector Off")
}

type SoundSystem struct{}

func (soundSystem *SoundSystem) On() {
	fmt.Println("SoundSystem On")
}

func (soundSystem *SoundSystem) Off() {
	fmt.Println("SoundSystem Off")
}

type HomeTheaterFacade struct {
	DvdPlayer   *DVDPlayer
	Projector   *Projector
	SoundSystem *SoundSystem
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{DvdPlayer: &DVDPlayer{},
		Projector:   &Projector{},
		SoundSystem: &SoundSystem{}}
}

func (htf *HomeTheaterFacade) WatchFilm() {
	fmt.Println("Starting watch film")
	htf.DvdPlayer.On()
	htf.Projector.On()
	htf.SoundSystem.On()
}

func (htf *HomeTheaterFacade) EndFilm() {
	fmt.Println("Ending watch film")
	htf.DvdPlayer.Off()
	htf.Projector.Off()
	htf.SoundSystem.Off()
}

func main() {
	theater := NewHomeTheaterFacade()
	theater.WatchFilm()
	fmt.Println()
	theater.EndFilm()
}
