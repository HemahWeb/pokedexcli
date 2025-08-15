package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Second)
	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)
	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key to be present in cache")
	}
	if string(got) != string(val) {
		t.Errorf("expected value %q, got %q", val, got)
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)
	key := "expire-key"
	val := []byte("expire-value")

	cache.Add(key, val)
	_, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key to be present in cache before expiration")
	}
	time.Sleep(150 * time.Millisecond)
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected key to be expired and not present in cache")
	}
}

func TestCacheDeleteExpired(t *testing.T) {
	cache := NewCache(50 * time.Millisecond)
	key := "delete-key"
	val := []byte("delete-value")

	cache.Add(key, val)
	time.Sleep(60 * time.Millisecond)
	cache.DeleteExpired()
	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key to be deleted after expiration and DeleteExpired call")
	}
}
