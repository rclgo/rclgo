package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rclgo/rclgo"
	"github.com/rclgo/stdmsgs"
)

func main() {
	node, err := rclgo.NewNode("GoSubscriber", "MyNamespace")
	if err != nil {
		panic(err)
	}

	var msg stdmsgs.String
	msg.Init()
	err = node.CreateSubscription(msg, "/cat_topic", myCallback)
	if err != nil {
		panic(err)
	}

	pub, err := node.CreatePublisher(msg, "/dog_topic")
	if err != nil {
		panic(err)
	}

	go rclgo.Spin(&node)

	go func() {
		i := 1
		for {
			time.Sleep(1_500 * time.Millisecond)
			var msg = stdmsgs.NewString("Woof woof " + strconv.Itoa(i))
			i++
			fmt.Println("Publishing:", msg.DataAsString())
			pub.Publish(msg)
			msg.Destroy()
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// subscriber.subscriber.Fini(subscriber.node)
	errs := node.Fini()
	if len(errs) != 0 {
		panic(errs)
	}

	err = rclgo.Shutdown()
	if err != nil {
		panic(err)
	}
	fmt.Println("done")
}

func myCallback(msg rclgo.RosMessage) {
	fmt.Printf("Received msg: %v\n", msg)
	m, ok := msg.(stdmsgs.String)
	if !ok {
		fmt.Println("msg is not std msg")
	} else {
		fmt.Println(m.DataAsString())
		fmt.Println(m.MsgInfo)
		// m.Destroy() // ToDo: This causes a double-free heap corruption. Learn why
	}
}
