package adapters

import (
	"arquitecturaHexagonal/src/books/domain"
	"arquitecturaHexagonal/src/helpers"
	"context"
	"encoding/json"
	"log"
	"time"
	"github.com/rabbitmq/amqp091-go"
)

type 	Rabbit struct {
	conn *helpers.RabbitMQConn
}

// Constructor para Rabbit, manteniendo la conexión abierta
func NewRabbit() *Rabbit {
	conn, err := helpers.GetRabbitMQConn()
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
	}

	return &Rabbit{conn: conn}
}

// Método para enviar un mensaje a RabbitMQ
func (rabbit *Rabbit) PublishEvent(book domain.Book) error {
	// Declaración del exchange (esto solo debería hacerse una vez)
	err := rabbit.conn.Ch.ExchangeDeclare(
		"logs",   // Nombre del exchange
		"fanout", // Tipo de exchange
		true,     // Durable
		false,    // Auto-deleted
		false,    // Internal
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		log.Printf("Error al declarar el exchange: %v", err)
		return err
	}

	// Convertir el libro a JSON
	body, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error al serializar el mensaje: %v", err)
		return err
	}

	// Contexto con timeout para la publicación
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publicar el mensaje
	err = rabbit.conn.Ch.PublishWithContext(ctx,
		"logs", // Exchange
		"",     // Routing key (vacío en fanout)
		false,  // Mandatory
		false,  // Immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Error al publicar mensaje: %v", err)
		return err
	}

	log.Printf(" [x] Mensaje enviado: %s", body)
	return nil
}
