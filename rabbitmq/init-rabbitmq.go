package rabbitmq

func InitRabbitMQ() {
	if client == nil {
		client = NewClient()
		client.Start()
	}
}

func CLoseRabbitMQ() {
	if client != nil {
		client.Stop()
	}
}