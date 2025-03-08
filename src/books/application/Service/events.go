package service

import (
	repositories "arquitecturaHexagonal/src/books/application/Repositories"
	dominio "arquitecturaHexagonal/src/books/domain"
)

type Event struct {
	rabbit repositories.IEventPublisher
}

func NewEvent(rabbit repositories.IEventPublisher)*Event{
	return &Event{rabbit:rabbit}
}

func (event *Event) Execute(book dominio.Book)(err error){
	return event.rabbit.PublishEvent(book)
}