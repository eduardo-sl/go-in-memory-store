package main

import (
	"bytes"
	"fmt"

	resp "go-in-memory-store/resp"
)

const (
	CommandSET   = "SET"
	CommandGET   = "GET"
	CommandHELLO = "hello"
)

type Command interface {
}

type SetCommand struct {
	key, val []byte
}

type HelloCommand struct {
	value string
}

type GetCommand struct {
	key []byte
}

func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buf)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":" + v)
	}
	return buf.Bytes()
}
