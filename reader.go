package main

import stomp "github.com/go-stomp/stomp"
import "os"
import "fmt"

//Connect to ActiveMQ and listen for messages
func main() {
	conn, err := stomp.Dial("tcp", os.Getenv("AMQ_ADDRESS"))
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("/queue/test-1", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(msg)
	}

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Disconnect()
}
