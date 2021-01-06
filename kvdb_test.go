package main

import (
	"testing"

	kvdb "github.com/kvdb/src"
)

func TestSet(t *testing.T) {
	res := kvdb.Set("k1", "10")
	if res != true {
		t.Errorf("Error: Unable to add key to kvdb")
	}
}

func TestGet(t *testing.T) {
	_, found := kvdb.Get("k1")
	if found != true {
		t.Errorf("Failed to find the key")
	}
}

func TestIncr(t *testing.T) {
	res := kvdb.Incr("k1")
	if res != true {
		t.Errorf("Error: Unable to increment value")
	}
}

func TestIncrBy(t *testing.T) {
	res := kvdb.IncrBy("k1", 20)
	if res != true {
		t.Errorf("Error: Unable to increment value")
	}
}

func TestDel(t *testing.T) {
	res := kvdb.Del("k1")
	if res != true {
		t.Errorf("Key not found!")
	}
}
