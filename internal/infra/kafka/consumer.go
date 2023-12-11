package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return err
	}

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		consumer.Close() // Feche o consumidor se falhar ao se inscrever nos tópicos
		return err
	}

	defer consumer.Close() // Certifique-se de fechar o consumidor quando a função retornar

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			return err // Retornar erro se ocorrer um problema ao ler a mensagem
		}

		msgChan <- msg
	}
}
