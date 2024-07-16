package tokenBucket

import (
	"testing"
	"time"
)

func TestTokenBucket_Take(t *testing.T) {
	bucket := NewTokenBucket(5, 5)

	t.Logf("test second token bucket start ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test token bucket allow take fail")
		}
	}

	t.Logf("test token bucket forbit ,size %d ", bucket.GetSize())
	if bucket.Take() {
		t.Fatalf("test token bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 1001)

	t.Logf("test token bucket again allow ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test token bucket allow take fail")
		}
	}
	if bucket.Take() {
		t.Fatalf("test token bucket forbit take fail")
	}

	t.Logf("test token bucket succes ,size %d ", bucket.GetSize())
}

func TestWithMilliseconds(t *testing.T) {
	bucket := NewTokenBucket(5, 5, WithMilliseconds(100))

	t.Logf("test milliseconds token bucket start ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test milliseconds token bucket allow take fail")
		}
	}

	t.Logf("test milliseconds token bucket forbit ,size %d ", bucket.GetSize())
	if bucket.Take() {
		t.Fatalf("test milliseconds token bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 101)

	t.Logf("test milliseconds token bucket again allow ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test milliseconds token bucket allow take fail")
		}
	}
	if bucket.Take() {
		t.Fatalf("test milliseconds token bucket forbit take fail")
	}

	t.Logf("test milliseconds token bucket succes ,size %d ", bucket.GetSize())
}
