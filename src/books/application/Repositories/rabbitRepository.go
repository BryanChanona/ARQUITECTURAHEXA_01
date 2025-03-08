package repositories

import "arquitecturaHexagonal/src/books/domain"

type IEventPublisher interface {
	PublishEvent(book domain.Book) error
}