package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	if cache.cache == nil {
		t.Error("NewCache did not initialize cache map")
	}
}

func TestAdd(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	cache.Add("testKey", []byte("testValue"))

	if _, ok := cache.cache["testKey"]; !ok {
		t.Error("Add did not add the key to the cache")
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	cache.Add("testKey", []byte("testValue"))

	val, ok := cache.Get("testKey")
	if !ok {
		t.Error("Get did not find the key in the cache")
	}

	if string(val) != "testValue" {
		t.Errorf("Get returned wrong value: got %v want %v", string(val), "testValue")
	}
}

func TestGetNonExistentKey(t *testing.T) {
	cache := NewCache(time.Minute * 5)

	_, ok := cache.Get("nonExistentKey")
	if ok {
		t.Error("Get found a key that was not added to the cache")
	}
}

func TestCacheEntryTimestamp(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	cache.Add("testKey", []byte("testValue"))

	entry, ok := cache.cache["testKey"]
	if !ok {
		t.Error("Add did not add the key to the cache")
	}

	if time.Since(entry.createdAt) > time.Second {
		t.Error("Cache entry timestamp is not recent")
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(time.Millisecond * 5)
	cache.Add("testKey1", []byte("testValue1"))
	cache.Add("testKey2", []byte("testValue2"))

	// Artificially age the first entry
	entry := cache.cache["testKey1"]
	entry.createdAt = time.Now().Add(-10 * time.Minute)
	cache.cache["testKey1"] = entry

	// Reap entries older than 5 minutes
	cache.reap(5 * time.Minute)

	// testKey1 should be reaped
	if _, ok := cache.Get("testKey1"); ok {
		t.Error("reap did not remove the expected key from the cache")
	}

	// testKey2 should still be there
	if _, ok := cache.Get("testKey2"); !ok {
		t.Error("reap removed a key that was not expected to be removed")
	}
}

func TestReapEmptyCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 5)

	// Reap entries older than 5 minutes
	cache.reap(5 * time.Minute)

	// Cache should still be empty
	if len(cache.cache) != 0 {
		t.Error("reap added entries to an empty cache")
	}
}
