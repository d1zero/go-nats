package main

import (
	"context"
	"encoding/json"
	"fmt"
	gonats "github.com/nats-io/nats.go"
	"go-nats/internal/dto"
	"log"
	"time"
)

const (
	logsSubject = "logs"
)

func sendTimeLog(conn *gonats.Conn, ctx context.Context, p dto.Logs) error {
	msg, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return conn.Publish(logsSubject, msg)
}

func main() {
	natsConn, err := gonats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer natsConn.Close()

	for {
		err := sendTimeLog(natsConn, context.Background(), dto.Logs{Text: time.Now().String()})
		if err != nil {
			panic(err)
		}
		fmt.Println("sended")
		time.Sleep(1 * time.Second)
	}
}
