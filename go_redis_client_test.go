package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

var server *Server

func setup() {
	listenAddr := ":5001"
	if server == nil {
		server = NewServer(Config{
			ListenAddr: listenAddr,
		})
		go func() {
			log.Fatal(server.Start())
		}()
		time.Sleep(time.Second)

		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:5001",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
}

func TestSetKey(t *testing.T) {
	setup()
	key := "foo"
	val := "bar"

	// set key
	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	setup()
	key := "foo"

	_, err := rdb.Get(context.Background(), key+"bar").Result()
	if err.Error() != "redis: nil" {
		t.Fatal(err)
	}
}

func TestGetKey(t *testing.T) {
	setup()
	key := "foo"
	val := "bar"

	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	newVal, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}

	if newVal != val {
		t.Fatalf("expected %s but got %s", val, newVal)
	}
}

func TestKeyExists(t *testing.T) {
	setup()
	key := "foo"
	val := "bar"

	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	existsVal, err := rdb.Exists(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}
	if existsVal != 1 {
		t.Fatalf("expected 1 but got %d", existsVal)
	}
}

func TestDeleteKeys(t *testing.T) {
	setup()
	key := "foo"
	val := "bar"

	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}
	if err := rdb.Set(context.Background(), key+"2", val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	delResult, err := rdb.Del(context.Background(), key, key+"1", key+"2").Result()
	if err != nil {
		t.Fatal(err)
	}
	if delResult != 2 {
		t.Fatalf("expected 2 but got %d", delResult)
	}

	existsVal, err := rdb.Exists(context.Background(), key).Result()
	if err != nil {
		t.Fatal(err)
	}
	if existsVal != 0 {
		t.Fatalf("expected 0 but got %d", existsVal)
	}
}

func TestExpireKey(t *testing.T) {
	setup()
	key := "foo"
	val := "bar"

	if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
		t.Fatal(err)
	}

	expirateSuccess, err := rdb.Expire(context.Background(), key, 10*time.Second).Result()
	if err != nil {
		t.Fatal(err)
	}
	if !expirateSuccess {
		t.Fatalf("expected true but got %t", expirateSuccess)
	}
}
