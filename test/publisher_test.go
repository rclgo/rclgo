package rclgo_test

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/rclgo/rclgo"
	"github.com/rclgo/stdmsgs"
)

// func BenchmarkFoo(b *testing.B) {

// 	for n := 0; n < b.N; n++ {

// 		// Initialization
// 		rcl.Init()
// 		rcl.Shutdown()

// 	}

// }

// func BenchmarkPublisher(b *testing.B) {

// 	for n := 0; n < b.N; n++ {

// 		// Initialization
// 		// rcl.Init()
// 		// myNode := node.GetZeroInitializedNode()
// 		// myNodeOpts := node.GetNodeDefaultOptions()

// 		// // fmt.Printf("Creating the node! \n")
// 		// node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 		// //Create the publisher
// 		// myPub := GetZeroInitializedPublisher()
// 		// myPubOpts := GetPublisherDefaultOptions()

// 		//Create the msg type
// 		var myMsg stdmsgs.String
// 		myMsg.Init()

// 		// fmt.Printf("Creating the publisher! \n")
// 		//Initializing the publisher
// 		// PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 		// index := 0

// 		// //Update my msg
// 		// myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
// 		// //Publish the message
// 		// Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		// fmt.Printf("Shutting down!! \n")

// 		myMsg.Destroy()
// 		// PublisherFini(myPub, myNode)
// 		// node.NodeFini(myNode)
// 		// rcl.Shutdown()
// 	}

// }

func TestPublisherStringMsg(t *testing.T) {

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

	myNode, err := rclgo.NewNode("GoPublisher", "")
	if err != nil {
		t.Fatalf("rclgo.NewNode(): %s\n", err)
	}

	// Create the msg type
	var myMsg stdmsgs.String
	myMsg.Init()

	// Create the publisher
	myPub, err := myNode.CreatePublisher(myMsg, "/myGoTopic")
	if err != nil {
		t.Fatalf("myNode.CreatePublisher(): %s\n", err)
	}

	index := 0
loop:
	for {
		//Update my msg
		myMsg.SetText("Greetings from GO! #" + strconv.Itoa(index))
		//Publish the message
		err := myPub.Publish(myMsg)
		if err != nil {
			t.Fatalf("myPub.Publish(): %s\n", err)
		} else {
			fmt.Printf("(Publisher) Published: %s\n", myMsg.DataAsString())
		}
		time.Sleep(500 * time.Millisecond)
		index++

		//Loop breaker
		select {
		case <-sigs:
			fmt.Println("Got shutdown, exiting")
			// Break out of the outer for statement and end the program
			break loop
		case <-msg:

		}
	}

	fmt.Printf("Shutting down!! \n")

	err = myMsg.Destroy()
	if err != nil {
		t.Fatalf("myMsg.Destroy(): %s\n", err)
	}

	// Publisher.Fini() called within myNode.Fini()
	// err = myPub.Fini(myNode)
	// if err != nil {
	// 	t.Fatalf("myPub.Fini(): %s\n", err)
	// }

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
		t.Fatalf("rclgo.Shutdown(): %s\n", err)
	}
}

// func TestPublisherInt8Msg(t *testing.T) {

// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	msg := make(chan string, 1)
// 	go func() {
// 		// Receive input in a loop
// 		for {
// 			var s string
// 			fmt.Scan(&s)
// 			// Send what we read over the channel
// 			msg <- s
// 		}
// 	}()

// 	// Initialization
// 	rcl.Init()
// 	myNode := node.GetZeroInitializedNode()
// 	myNodeOpts := node.GetNodeDefaultOptions()

// 	fmt.Printf("Creating the node! \n")
// 	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 	//Create the publisher
// 	myPub := GetZeroInitializedPublisher()
// 	myPubOpts := GetPublisherDefaultOptions()

// 	//Create the msg type
// 	var myMsg stdmsgs.Int8
// 	myMsg.Init()
// 	myMsg.SetInt8(32)

// 	fmt.Printf("Creating the publisher! \n")
// 	//Initializing the publisher
// 	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 	index := 0
// loop:
// 	for {
// 		//Update my msg
// 		myMsg.SetInt8(myMsg.GetInt8() + 1)

// 		//Publish the message
// 		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		if retRCL == cwrap.Ok {
// 			var value string
// 			fmt.Sprintf("%v", value)

// 			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetInt8())
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 		index++

// 		//Loop breaker
// 		select {
// 		case <-sigs:
// 			fmt.Println("Got shutdown, exiting")
// 			// Break out of the outer for statement and end the program
// 			break loop
// 		case <-msg:

// 		}
// 	}

// 	fmt.Printf("Shutting down!! \n")

// 	myMsg.Destroy()
// 	PublisherFini(myPub, myNode)
// 	node.NodeFini(myNode)
// 	rcl.Shutdown()

// }

// func TestPublisherFloat64Msg(t *testing.T) {

// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	msg := make(chan string, 1)
// 	go func() {
// 		// Receive input in a loop
// 		for {
// 			var s string
// 			fmt.Scan(&s)
// 			// Send what we read over the channel
// 			msg <- s
// 		}
// 	}()

// 	// Initialization
// 	rcl.Init()
// 	myNode := node.GetZeroInitializedNode()
// 	myNodeOpts := node.GetNodeDefaultOptions()

// 	fmt.Printf("Creating the node! \n")
// 	node.NodeInit(myNode, "GoPublisher", "", myNodeOpts)

// 	//Create the publisher
// 	myPub := GetZeroInitializedPublisher()
// 	myPubOpts := GetPublisherDefaultOptions()

// 	//Create the msg type
// 	var myMsg stdmsgs.Float64
// 	myMsg.Init()
// 	myMsg.SetFloat64(32.0)

// 	fmt.Printf("Creating the publisher! \n")
// 	//Initializing the publisher
// 	PublisherInit(myPub, myPubOpts, myNode, "/myGoTopic", myMsg.GetMessage())

// 	index := 0
// loop:
// 	for {
// 		//Update my msg
// 		myMsg.SetFloat64(myMsg.GetFloat64() + 0.25)

// 		//Publish the message
// 		retRCL := Publish(myPub, myMsg.GetMessage(), myMsg.GetData())

// 		if retRCL == cwrap.Ok {
// 			var value string
// 			fmt.Sprintf("%v", value)

// 			fmt.Printf("(Publisher) Published: %v\n", myMsg.GetFloat64())
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 		index++

// 		//Loop breaker
// 		select {
// 		case <-sigs:
// 			fmt.Println("Got shutdown, exiting")
// 			// Break out of the outer for statement and end the program
// 			break loop
// 		case <-msg:

// 		}
// 	}

// 	fmt.Printf("Shutting down!! \n")

// 	myMsg.Destroy()
// 	PublisherFini(myPub, myNode)
// 	node.NodeFini(myNode)
// 	rcl.Shutdown()

// }
