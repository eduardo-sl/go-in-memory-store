package main

// Ucomment the following code to run the test runnig the server in the machine

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	"testing"
// 	"time"

// 	"go-in-memory-store/resp"
// )

// type Client struct {
// 	addr string
// 	conn net.Conn
// }

// func New(addr string) (*Client, error) {
// 	conn, err := net.Dial("tcp", addr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Client{
// 		addr: addr,
// 		conn: conn,
// 	}, nil
// }

// func (c *Client) Set(ctx context.Context, key string, val string) error {

// 	var buf bytes.Buffer
// 	wr := resp.NewWriter(&buf)
// 	wr.WriteArray([]resp.Value{
// 		resp.StringValue("set"),
// 		resp.StringValue(key),
// 		resp.StringValue(val),
// 	})

// 	_, err := c.conn.Write(buf.Bytes())
// 	return err
// }

// func (c *Client) Get(ctx context.Context, key string) (string, error) {
// 	var buf bytes.Buffer
// 	wr := resp.NewWriter(&buf)
// 	wr.WriteArray([]resp.Value{
// 		resp.StringValue("get"),
// 		resp.StringValue(key),
// 	})

// 	_, err := c.conn.Write(buf.Bytes())
// 	if err != nil {
// 		return "", err

// 	}
// 	b := make([]byte, 1024)
// 	n, err := c.conn.Read(b)

// 	return string(b[:n]), err
// }

// func (c *Client) Close() error {
// 	return c.conn.Close()
// }

// func TestNewClient(t *testing.T) {
// 	c, err := New("localhost:5001")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()
// 	time.Sleep(time.Second)
// 	if err := c.Set(context.Background(), "foo", "bar"); err != nil {
// 		log.Fatal(err)
// 	}

// 	val, err := c.Get(context.Background(), "foo")
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	fmt.Println("GET =>", val)
// }
