package rclgo_test

import (
	"testing"
	"time"

	"github.com/rclgo/rclgo"
)

func TestNodeCreation(t *testing.T) {
	// Initialization
	myNode, err := rclgo.NewNode("fakeNameForNode", "")
	if err != nil {
		t.Fatalf("rclgo.NewNode(): %s\n", err)
	}

	time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee

	errs := myNode.Fini()
	if len(errs) != 0 {
		var errMsg string = "myNode.Fini():\n"
		for _, err := range errs {
			errMsg += err.Error() + "\n"
		}
		t.Fatalf(errMsg)
	}

	err = rclgo.Shutdown()
	if err != nil {
		t.Fatalf("rclgo.Shutdown(): %s\n", err)
	}
}
