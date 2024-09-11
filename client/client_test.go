package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewClientRedisClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Set(context.Background(), "key", "value", 0).Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("WE ARE HERE")
}

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
