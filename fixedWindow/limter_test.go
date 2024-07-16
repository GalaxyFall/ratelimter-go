package fixedWindow

import (
	"testing"
	"time"
)

func TestFixedWindow_Take(t *testing.T) {
	window := NewFixedWindow(5)

	t.Logf("test second fixed window start ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test fixed window allow take fail")
		}
	}

	t.Logf("test fixed window forbit ,size %d ", window.GetSize())
	if window.Take() {
		t.Fatalf("test fixed window forbit take fail")
	}

	time.Sleep(time.Millisecond * 1001)

	t.Logf("test fixed window again allow ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test fixed window allow take fail")
		}
	}
	if window.Take() {
		t.Fatalf("test fixed window forbit take fail")
	}

	t.Logf("test fixed window succes ,size %d ", window.GetSize())
}

func TestWithMilliseconds(t *testing.T) {
	window := NewFixedWindow(5, WithMilliseconds(100))

	t.Logf("test milliseconds fixed window start ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test milliseconds fixed window allow take fail")
		}
	}

	t.Logf("test milliseconds fixed window bucket forbit ,size %d ", window.GetSize())
	if window.Take() {
		t.Fatalf("test milliseconds fixed window forbit take fail")
	}

	time.Sleep(time.Millisecond * 101)

	t.Logf("test milliseconds fixed window bucket again allow ,size %d ", window.GetSize())
	for i := 0; i < 5; i++ {
		if !window.Take() {
			t.Fatalf("test milliseconds fixed window allow take fail")
		}
	}
	if window.Take() {
		t.Fatalf("test milliseconds fixed window forbit take fail")
	}

	t.Logf("test milliseconds fixed window succes ,size %d ", window.GetSize())
}
