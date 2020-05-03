package rclgo_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/rclgo/rclgo"
	"github.com/rclgo/stdmsgs"
)

func TestSubscription(t *testing.T) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()

	// Initialization
	myNode, err := rclgo.NewNode("GoSubscriber", "")
	if err != nil {
		t.Fatalf("rclgo.NewNode(): %s", err)
	}

	// Creating the msg type
	var myMsg stdmsgs.String
	myMsg.Init()

	// Create the subscriptor
	err = myNode.CreateSubscription(myMsg, "/myGoTopic", func(rm rclgo.RosMessage) { callback(t, rm) })
	if err != nil {
		t.Fatalf("myNode.CreateSubscription(): %s", err)
	}

loop:
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			break loop
		case <-msg:
		}
	}

	fmt.Printf("Shutting down!! \n")

	err = myMsg.Destroy()
	if err != nil {
		t.Fatalf("myMsg.Destroy(): %s", err)
	}

	errs := myNode.Fini()
	if len(errs) != 0 {
		var s = "myNode.Fini():\n"
		for _, err := range errs {
			s += err.Error() + "\n"
		}
		t.Fatalf(s)
	}

	err = rclgo.Shutdown()
	if err != nil {
		t.Fatalf("rclgo.Shutdown(): %s", err)
	}
}

func callback(t *testing.T, msg rclgo.RosMessage) {
	str, ok := msg.(stdmsgs.String)
	if !ok {
		t.Fatalf("callback msg could not be asserted as stdmsgs.String")
	}
	fmt.Println("Received:", str.DataAsString())
}
