package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"go-nats/internal/dto"
	"os"
	"os/signal"
	"syscall"
)

func handler(msg *nats.Msg) {
	var p dto.Logs
	if err := json.Unmarshal(msg.Data, &p); err != nil {
		fmt.Printf("unmarshal json: %s\n", err.Error())
		return
	}

	fmt.Println(p.Text)
}

func main() {
	natsConn, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		panic(err)
	}
	defer natsConn.Close()
	fmt.Println("connected to nats")

	logsSub, err := natsConn.Subscribe("logs", handler)
	if err != nil {
		fmt.Errorf("subscribe to logs subject: %w", err)
	}

	defer func() {
		err = logsSub.Unsubscribe()
		if err != nil {
			fmt.Errorf("unsubscribe logs subject: %s", err.Error())
		}
	}()

	fmt.Println("Application has started")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	fmt.Println("Application has been shut down")
}
