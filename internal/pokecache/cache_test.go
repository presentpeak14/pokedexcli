package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	c := NewCache(5 * time.Second)
	c.Add("https://example.com", []byte("data"))
	val, ok := c.Get("https://example.com")
	if !ok {
		t.Fatal("expected key to exist in cache")
	}
	if string(val) != "data" {
		t.Errorf("expected 'data', got %q", string(val))
	}
}

func TestGetMissing(t *testing.T) {
	c := NewCache(5 * time.Second)
	_, ok := c.Get("https://nothere.com")
	if ok {
		t.Fatal("expected key to not exist in cache")
	}
}

func TestReap(t *testing.T) {
	interval := 50 * time.Millisecond
	c := NewCache(interval)
	c.Add("https://example.com", []byte("data"))

	time.Sleep(interval * 3)

	_, ok := c.Get("https://example.com")
	if ok {
		t.Fatal("expected entry to have been reaped")
	}
}
