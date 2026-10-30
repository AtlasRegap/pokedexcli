package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache := NewCache(5 * time.Second)
		cache.Add(c.key, c.val)

		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %q", c.key)
			continue
		}
		if string(val) != string(c.val) {
			t.Errorf("key %q: expected %q, got %q", c.key, c.val, val)
		}
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("not-there")
	if ok {
		t.Errorf("expected missing key to return false")
	}
}

func TestReapLoop(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key right after adding")
		return
	}

	time.Sleep(3 * interval)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected key to be reaped")
	}
}
