package mqtt

import (
	"errors"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	client mqtt.Client
}

func NewMqttService(clientID, brokerIp string, port int) (*MqttService, error) {
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerIp, port))

	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, errors.New("Error to get token")
	}

	return &MqttService{
		client: client,
	}, nil
}

func (m *MqttService) Publish(topic, message string) {
	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message

	token := m.client.Publish(topic, 0, false, message)
	token.Wait()

}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}
