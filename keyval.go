package main

import "sync"

type KV struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewKV() *KV {
	return &KV{
		data: make(map[string][]byte),
	}
}

func (kv *KV) Set(key, val []byte) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[string(key)] = []byte(val)
	return nil
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[string(key)]
	return val, ok
}

func (kv *KV) Del(key []byte) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	_, ok := kv.data[string(key)]
	if !ok {
		return false
	}
	delete(kv.data, string(key))
	return true
}
