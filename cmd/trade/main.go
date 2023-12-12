package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/infra/kafka"
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/market/dto"
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/market/entity"
	"github.com/lucasfiduniv/HomeBroker-microservices-goLang/internal/market/transformer"
)

func main() {
	orderIn := make(chan *entity.Order)
	orderOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgCha := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.server":  "host.docker.internal:9094",
		"group.id":          "mygroup",
		"auto.offset.reset": "earliest",
	}
	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgCha)

	book := entity.NewBook(orderIn, orderOut, wg)
	go book.Trade()

	go func() {
		for msg := range kafkaMsgChan {
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformerInput(tradeInput)
			orderIn <- order
		}
	}()
	for res := range ordersOut {
		output := transformer.TranformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "output")
	}
}
