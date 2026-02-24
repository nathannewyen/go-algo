package lru

import (
	"fmt"
	"testing"
)

func TestNewCache(t *testing.T) {
	lruCache := NewCache(3)

	if lruCache.Size() != 0 {
		t.Errorf("expected size 0, got %d", lruCache.Size())
	}
}

func TestPutAndGet(t *testing.T) {
	lruCache := NewCache(3)

	lruCache.Put("a", 1)
	lruCache.Put("b", 2)

	value, found := lruCache.Get("a")
	if !found {
		t.Error("expected to find key 'a'")
	}
	if value != 1 {
		t.Errorf("expected value 1, got %d", value)
	}
}

func TestGetMissing(t *testing.T) {
	lruCache := NewCache(3)

	_, found := lruCache.Get("nonexistent")
	if found {
		t.Error("expected get on missing key to return false")
	}
}

func TestEviction(t *testing.T) {
	lruCache := NewCache(2)

	lruCache.Put("a", 1)
	lruCache.Put("b", 2)
	// Cache is full: {b, a}

	// Adding "c" should evict "a" (least recently used)
	lruCache.Put("c", 3)

	_, found := lruCache.Get("a")
	if found {
		t.Error("expected key 'a' to be evicted")
	}

	value, found := lruCache.Get("b")
	if !found {
		t.Error("expected key 'b' to still exist")
	}
	if value != 2 {
		t.Errorf("expected value 2, got %d", value)
	}

	value, found = lruCache.Get("c")
	if !found {
		t.Error("expected key 'c' to exist")
	}
	if value != 3 {
		t.Errorf("expected value 3, got %d", value)
	}
}

func TestGetUpdatesRecency(t *testing.T) {
	lruCache := NewCache(2)

	lruCache.Put("a", 1)
	lruCache.Put("b", 2)

	// Access "a" to make it more recently used than "b"
	lruCache.Get("a")

	// Adding "c" should now evict "b" (least recently used)
	lruCache.Put("c", 3)

	_, found := lruCache.Get("b")
	if found {
		t.Error("expected key 'b' to be evicted after 'a' was accessed")
	}

	_, found = lruCache.Get("a")
	if !found {
		t.Error("expected key 'a' to still exist")
	}
}

func TestPutUpdateExisting(t *testing.T) {
	lruCache := NewCache(2)

	lruCache.Put("a", 1)
	lruCache.Put("a", 100)

	// Updating should not increase the size
	if lruCache.Size() != 1 {
		t.Errorf("expected size 1 after update, got %d", lruCache.Size())
	}

	value, found := lruCache.Get("a")
	if !found {
		t.Error("expected to find key 'a'")
	}
	if value != 100 {
		t.Errorf("expected updated value 100, got %d", value)
	}
}

func TestCapacityOne(t *testing.T) {
	lruCache := NewCache(1)

	lruCache.Put("a", 1)
	lruCache.Put("b", 2)

	_, found := lruCache.Get("a")
	if found {
		t.Error("expected key 'a' to be evicted with capacity 1")
	}

	value, found := lruCache.Get("b")
	if !found {
		t.Error("expected key 'b' to exist")
	}
	if value != 2 {
		t.Errorf("expected value 2, got %d", value)
	}
}

// BenchmarkPut measures the performance of cache insertions with evictions.
func BenchmarkPut(b *testing.B) {
	lruCache := NewCache(1000)

	for iterationIndex := 0; iterationIndex < b.N; iterationIndex++ {
		key := fmt.Sprintf("key-%d", iterationIndex)
		lruCache.Put(key, iterationIndex)
	}
}

// BenchmarkGet measures the performance of cache lookups.
func BenchmarkGet(b *testing.B) {
	lruCache := NewCache(1000)

	// Pre-populate the cache
	for setupIndex := 0; setupIndex < 1000; setupIndex++ {
		key := fmt.Sprintf("key-%d", setupIndex)
		lruCache.Put(key, setupIndex)
	}

	b.ResetTimer()

	for iterationIndex := 0; iterationIndex < b.N; iterationIndex++ {
		key := fmt.Sprintf("key-%d", iterationIndex%1000)
		lruCache.Get(key)
	}
}
