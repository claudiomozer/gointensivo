package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/claudiomozer/taxas/internal/infra/database"
	"github.com/claudiomozer/taxas/internal/usecase"
	"github.com/claudiomozer/taxas/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"

	go kafka.Consume(topics, servers, msgChanKafka)
	kafkaWorker(msgChanKafka, usecase)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	for msg := range msgChan {
		var OrderInputDto usecase.OrderInputDto
		err := json.Unmarshal(msg.Value, &OrderInputDto)

		if err != nil {
			panic(err)
		}

		outputDto, err := uc.Execute(OrderInputDto)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Kafka has processed order %v\n", outputDto)
	}
}
