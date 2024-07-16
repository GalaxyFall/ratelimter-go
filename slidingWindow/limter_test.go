package slidingWindow

import (
	"testing"
	"time"
)

func TestSlidingWindow_Take(t *testing.T) {
	window := NewSlidingWindow(5)

	t.Logf("test second sliding window start ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test sliding window bucket allow take fail")
		}
	}

	t.Logf("test sliding window bucket forbit ,size %d ", window.GetSize())
	if window.Take() {
		t.Fatalf("test sliding window bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 1001)

	t.Logf("test sliding window bucket again allow ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test sliding window bucket allow take fail")
		}
	}
	if window.Take() {
		t.Fatalf("test sliding window bucket forbit take fail")
	}

	t.Logf("test sliding window bucket succes ,size %d ", window.GetSize())
}

func TestWithMilliseconds(t *testing.T) {
	window := NewSlidingWindow(5, WithMilliseconds(100))

	t.Logf("test milliseconds sliding window start ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test milliseconds sliding window bucket allow take fail")
		}
	}

	t.Logf("test milliseconds sliding window bucket forbit ,size %d ", window.GetSize())
	if window.Take() {
		t.Fatalf("test milliseconds sliding window bucket forbit take fail")
	}

	time.Sleep(time.Millisecond * 101)

	t.Logf("test milliseconds sliding window bucket again allow ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test milliseconds sliding window bucket allow take fail")
		}
	}
	if window.Take() {
		t.Fatalf("test milliseconds sliding window bucket forbit take fail")
	}

	t.Logf("test milliseconds sliding window bucket succes ,size %d ", window.GetSize())
}
