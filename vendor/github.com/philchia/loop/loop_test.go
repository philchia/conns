package loop

import (
	"fmt"
	"testing"
)

func TestManyThings(t *testing.T) {
	l := New(20)
	if l == nil {
		t.Errorf("New got a nil value")
	}

	for i := 0; i < 20; i++ {
		t.Log("Push")
		if ok := l.Push(fmt.Sprintf("%d", i)); !ok {
			t.Errorf("Push failed")
		}
	}

	for i := 0; i < 10; i++ {
		if ok := l.Push(fmt.Sprintf("%d", i)); ok {
			t.Errorf("Push should fail after full")
		}
	}

	for i := 0; i < 20; i++ {
		if obj := l.Pop(); obj == nil || obj.(string) != fmt.Sprintf("%d", i) {
			t.Errorf("Pop got %v want %s", obj, fmt.Sprintf("%d", i))
		}
	}

	for i := 0; i < 20; i++ {
		if obj := l.Pop(); obj != nil {
			t.Errorf("Pop got %v want nil", obj)
		}
	}
}
