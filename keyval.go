package main

import (
	"sync"
)

type Data struct {
	val []byte
	ttl int64
}

type KV struct {
	mu   sync.RWMutex
	data map[string]Data
}

func NewKV() *KV {
	return &KV{
		data: make(map[string]Data),
	}
}

func (kv *KV) Set(key, val []byte) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[string(key)] = Data{val: []byte(val), ttl: -1}
	return nil
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[string(key)]
	return val.val, ok
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

func (kv *KV) Expire(key []byte, exp int64) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[string(key)]
	if !ok {
		return false
	}
	kv.data[string(key)] = Data{val: []byte(val.val), ttl: exp}
	return true
}
