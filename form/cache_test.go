package form

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache()
	f := New("test", "test")
	id := "foo"

	expires := time.Now().Add(time.Second)
	if err := c.Set(id, f, expires); err != nil {
		t.Errorf("Failed to set cache: %s", err)
	}

	f2, err := c.Get(id)
	if err != nil {
		t.Errorf("Failed to get cached record: %s", err)
	}

	if f2.Name != f.Name {
		t.Errorf("Expected %q, got %q", f.Name, f2.Name)
	}

	c.Remove(id)
	if _, err := c.Get(id); err != FormNotFound {
		t.Errorf("Expected entry to be removed, but it's here.")
	}

}
