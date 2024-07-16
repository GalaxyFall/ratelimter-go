package leakyBucket

import (
	"testing"
	"time"
)

func TestLeakyBucket_Take(t *testing.T) {
	bucket := NewLeakyBucket(5, 5)

	t.Logf("test second leaky bucket start ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test leaky bucket allow take fail")
		}
	}

	t.Logf("test leaky bucket forbit ,size %d ", bucket.GetSize())
	if bucket.Take() {
		t.Fatalf("test leaky bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 1001)

	t.Logf("test leaky bucket again allow ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test leaky bucket allow take fail")
		}
	}

	if bucket.Take() {
		t.Fatalf("test leaky bucket forbit take fail")
	}

	t.Logf("test leaky bucket succes ,size %d ", bucket.GetSize())
}

func TestWithMilliseconds(t *testing.T) {

	bucket := NewLeakyBucket(5, 5, WithMilliseconds(100))

	t.Logf("test milliseconds leaky bucket start ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test milliseconds leaky bucket allow take fail")
		}
	}

	t.Logf("test milliseconds leaky bucket forbit ,size %d ", bucket.GetSize())
	if bucket.Take() {
		t.Fatalf("test milliseconds leaky bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 101)

	t.Logf("test milliseconds leaky bucket again allow ,size %d ", bucket.GetSize())
	for i := 0; i < 5; i++ {
		if !bucket.Take() {
			t.Fatalf("test milliseconds leaky bucket allow take fail")
		}
	}

	if bucket.Take() {
		t.Fatalf("test leaky bucket forbit take fail")
	}

	t.Logf("test milliseconds leaky bucket succes ,size %d ", bucket.GetSize())
}
