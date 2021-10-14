package messaging

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type MessageConnection struct {
	Client interface{}
}

type Message struct {
	Subject string
	Data    []byte
}

var ChMessage chan Message

func GetConnection(conStr string) (Conn *MessageConnection, err error) {
	nc, err := nats.Connect(conStr)
	if err != nil {
		return nil, err
	}
	return &MessageConnection{Client: nc}, nil
}

func Disconnect(Conn *MessageConnection) {
	Conn.Client.(*nats.Conn).Close()
	close(ChMessage)
}

func Publish(Conn *MessageConnection) {
	for chData := range ChMessage {
		//	fmt.Println(chData)
		err := Conn.Client.(*nats.Conn).Publish(chData.Subject, chData.Data)
		if err != nil {
			fmt.Println("--->>>>>>>>>>>>", err)
		}
	}
}

func Subscribe(Conn *MessageConnection) {
	Conn.Client.(*nats.Conn).Subscribe("person.create", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

}

// init is to initialize when package is first time caleld.
func Init(Conn *MessageConnection) {
	ChMessage = make(chan Message)
	go Publish(Conn)
	go Subscribe(Conn)
}
