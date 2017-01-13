package mongo

import (
	"testing"
	"time"

	"github.com/mataharimall/micro-api/config"
)

func init() {
	config.Init()
}

func TestConnection(t *testing.T) {
	conn := New("test")
	if conn == nil {
		t.Errorf("Could not access mongo db.")
	}
}

func TestExpiration(t *testing.T) {
	store := New("test")
	defer store.Flush()
	testValues := map[string]int{
		"v1": 3,
		"v2": 6,
	}

	if err := store.SetLifetime(time.Second * 1); err != nil {
		t.Skip("Set lifetime to all items is not supported")
	}

	if err := store.Add("v1", testValues["v1"]); err != nil {
		t.Errorf("Could not add value: %v", err)
	}
	if err := store.Add("v2", testValues["v2"]); err != nil {
		t.Errorf("Could not add value: %v", err)
	}
	var result int

	err := store.Get("v1", &result)
	if err != nil {
		t.Errorf("The value v1 was not stored: %v", err)
	}
	if err := store.Get("v2", &result); err != nil {
		t.Errorf("The value v2 was not stored: %v", err)
	}

	time.Sleep(time.Second * 3)

	err = store.Get("v1", &result)
	if err != nil {
		t.Errorf("The value v1 was not expired: %v", err)
	}
	err = store.Get("v2", &result)
	if err != nil {
		t.Errorf("The value v2 was not expired: %v", err)
	}

	err = store.Delete("v1")
	if err != nil {
		t.Errorf("The expired value v1 should not be removable: %v", err)
	}
	err = store.Delete("v2")
	if err != nil {
		t.Errorf("The expired value v2 should not be removable: %v", err)
	}
}
