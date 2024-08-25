package rabbitmq

import (
	"access-point/config"

	"github.com/gohugoio/hugo/publisher"
)

type Client struct {
	connectionManager *ConnectionManager
	consumerManager *consumerManager
	publisher *Publisher
}

var client *Client 

func GetClient() *Client {
	return client
}

func NewClient() *Client {
	conf := config.GetConfig()
	c := &Client {
		connectionManager: NewConnectionManager(

		),	
	}
}