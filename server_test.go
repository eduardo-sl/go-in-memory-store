package main

import (
	"context"
	"fmt"
	"go-in-memory-store/client"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestOfficialRedisClient(t *testing.T) {
	listenAddr := ":5001"
	server := NewServer(Config{
		ListenAddr: listenAddr,
	})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := "foo"
	val := "bar"

	// set key
	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	// get key
	_, err := rdb.Get(context.Background(), key+"bar").Result()
	if err.Error() != "redis: nil" {
		t.Fatal(err)
	}

	newVal, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}

	if newVal != val {
		t.Fatalf("expected %s but got %s", val, newVal)
	}

	// exists key
	existsVal, err := rdb.Exists(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}
	if existsVal != 1 {
		t.Fatalf("expected 1 but got %d", existsVal)
	}

	if err := rdb.Set(context.Background(), key+"2", val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	// del key
	delResult, err := rdb.Del(context.Background(), key, key+"1", key+"2").Result()
	if err != nil {
		t.Fatal(err)
	}
	if delResult != 2 {
		t.Fatalf("expected 2 but got %d", delResult)
	}

	existsVal, err = rdb.Exists(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}
	if existsVal != 0 {
		t.Fatalf("expected 0 but got %d", existsVal)
	}

	// expire key

	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	expirateSuccess, err := rdb.Expire(context.Background(), key, 10*time.Second).Result()
	if !expirateSuccess {
		t.Fatalf("expected true but got %t", expirateSuccess)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestRespWriteMap(t *testing.T) {
	in := map[string]string{
		"server":  "redis",
		"version": "6.0",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}

func TestServerWithMuiltClients(t *testing.T) {
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)
	nClients := 10
	wg := sync.WaitGroup{}
	wg.Add(nClients)
	for i := 0; i < nClients; i++ {
		go func(it int) {
			c, err := client.New("localhost:5001")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			key := fmt.Sprintf("client_foo_%d", i)
			value := fmt.Sprintf("client_bar_%d", i)
			if err := c.Set(context.TODO(), key, value); err != nil {
				log.Fatal(err)
			}
			val, err := c.Get(context.TODO(), key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("client %d got this value back => %s\n", it, val)

			wg.Done()
		}(i)
	}
	wg.Wait()

	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Fatalf("expected 0 peers but got %d", len(server.peers))
	}
}
