package main

import (
	"reflect"
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

func TestCompact(t *testing.T) {
	inpDB := map[string]string{"foo": "name", "bar": "game"}

	kvdb.Set("foo", "name")
	kvdb.Set("bar", "game")

	db := kvdb.Compact()
	if !reflect.DeepEqual(db, inpDB) {
		t.Errorf("Mismatch in getting map view!")
	}

}
