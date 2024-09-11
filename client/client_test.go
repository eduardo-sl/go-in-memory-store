package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	time.Sleep(time.Second)
	if err := c.Set(context.Background(), "foo", "bar"); err != nil {
		log.Fatal(err)
	}

	val, err := c.Get(context.Background(), "foo")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("GET =>", val)
}
