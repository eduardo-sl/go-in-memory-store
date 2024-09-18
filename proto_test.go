package main

import (
	"fmt"
	"testing"
)

func TestRespWriteMap(t *testing.T) {
	in := map[string]string{
		"server":  "redis",
		"version": "6.0",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}
